package mqtt

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"html/template"
	"net/url"
	"strconv"
	"strings"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/data/metadata"
	"github.com/project-flogo/core/support/log"
	"github.com/project-flogo/core/support/ssl"
)

var activityMd = activity.ToMetadata(&Settings{}, &Input{}, &Output{})

func init() {
	_ = activity.Register(&Activity{}, New)
}

// TokenType is a type of token
type TokenType int

const (
	// Literal is a literal token type
	Literal TokenType = iota
	// Substitution is a parameter substitution
	Substitution
)

// Token is a MQTT topic token
type Token struct {
	TokenType TokenType
	Token     string
}

// Topic is a parsed topic
type Topic []Token

// ParseTopic parses the topic
func ParseTopic(topic string) Topic {
	var parsed Topic
	parts, index := strings.Split(topic, "/"), 0
	for _, part := range parts {
		if strings.HasPrefix(part, ":") {
			token := strings.TrimPrefix(part, ":")
			if token == "" {
				token = strconv.Itoa(index)
				index++
			}
			parsed = append(parsed, Token{
				TokenType: Substitution,
				Token:     token,
			})
		} else {
			parsed = append(parsed, Token{
				TokenType: Literal,
				Token:     part,
			})
		}
	}
	return parsed
}

// String generates a string for the topic with params
func (t Topic) String(params map[string]string) string {
	output := strings.Builder{}
	for i, token := range t {
		if i > 0 {
			output.WriteString("/")
		}
		switch token.TokenType {
		case Literal:
			output.WriteString(token.Token)
		case Substitution:
			if value, ok := params[token.Token]; ok {
				output.WriteString(value)
			} else {
				output.WriteString(":")
				output.WriteString(token.Token)
			}
		}
	}
	return output.String()
}

func New(ctx activity.InitContext) (activity.Activity, error) {
	settings := &Settings{}

	err := metadata.MapToStruct(ctx.Settings(), settings, true)
	if err != nil {
		return nil, err
	}

	options := initClientOption(ctx.Logger(), settings)

	if strings.HasPrefix(settings.Broker, "ssl") {

		cfg := &ssl.Config{}

		if len(settings.SSLConfig) != 0 {
			err := cfg.FromMap(settings.SSLConfig)
			if err != nil {
				return nil, err
			}

			if _, set := settings.SSLConfig["skipVerify"]; !set {
				cfg.SkipVerify = true
			}
			if _, set := settings.SSLConfig["useSystemCert"]; !set {
				cfg.UseSystemCert = true
			}
		} else {
			//using ssl but not configured, use defaults
			cfg.SkipVerify = true
			cfg.UseSystemCert = true
		}

		tlsConfig, err := ssl.NewClientTLSConfig(cfg)
		if err != nil {
			return nil, err
		}

		options.SetTLSConfig(tlsConfig)
	}

	mqttClient := mqtt.NewClient(options)
	fmt.Println(mqttClient)

	fmt.Println("connecting")
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println(token)
		return nil, token.Error()
	}

	fmt.Println("mqtt success")
	act := &Activity{
		client:   mqttClient,
		settings: settings,
		topic:    ParseTopic(settings.Topic),
	}
	return act, nil
}

type Activity struct {
	settings *Settings
	client   mqtt.Client
	topic    Topic
}

func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}

	err = ctx.GetInputObject(input)

	if err != nil {
		return true, err
	}

	//fmt.Println(input.Password)
	//connString = input.Password
	//deviceID = input.DeviceId
	topic := a.settings.Topic
	if params := input.TopicParams; len(params) > 0 {
		topic = a.topic.String(params)
	}
	if token := a.client.Publish(topic, byte(a.settings.Qos), a.settings.Retain, input.Message); token.Wait() && token.Error() != nil {
		ctx.Logger().Debugf("Error in publishing: %v", err)
		return true, token.Error()
	}

	ctx.Logger().Debugf("Published Message: %v", input.Message)

	return true, nil
}

func tryGetKeyByName(v url.Values, key string) string {
	if len(v[key]) == 0 {
		return ""
	}

	return strings.Replace(v[key][0], " ", "+", -1)
}

func fetchpassword(connString string) string {

	url, err := url.ParseQuery(connString)
	if err != nil {
		fmt.Println(err)
	}

	deviceID := "00:15:5d:01:6d:00"

	h := tryGetKeyByName(url, "HostName")
	kn := tryGetKeyByName(url, "SharedAccessKeyName")
	k := tryGetKeyByName(url, "SharedAccessKey")
	//d := tryGetKeyByName(url, "DeviceId")
	//fmt.Println(h)
	//fmt.Println(kn)
	//fmt.Println(k)
	//fmt.Println(deviceID)

	uri := fmt.Sprintf("%s/twins/%s?api-version=2018-06-30", h, deviceID)
	timestamp := time.Now().Unix() + int64(3600)
	encodedURI := template.URLQueryEscaper(uri)

	toSign := encodedURI + "\n" + strconv.FormatInt(timestamp, 10)

	binKey, _ := base64.StdEncoding.DecodeString(k)
	mac := hmac.New(sha256.New, []byte(binKey))
	mac.Write([]byte(toSign))

	encodedSignature := template.URLQueryEscaper(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	if kn != "" {
		return fmt.Sprintf("SharedAccessSignature sig=%s&se=%d&skn=%s&sr=%s", encodedSignature, timestamp, kn, encodedURI)
	}
	fmt.Sprintf("SharedAccessSignature sig=%s&se=%d&sr=%s", encodedSignature, timestamp, encodedURI)

	return fmt.Sprintf("SharedAccessSignature sig=%s&se=%d&sr=%s", encodedSignature, timestamp, encodedURI)
}

func initClientOption(logger log.Logger, settings *Settings) *mqtt.ClientOptions {

	fmt.Println(settings.Connstring)
	password := fetchpassword(settings.Connstring)
	fmt.Println(settings.Broker)
	fmt.Println(settings.Id)
	fmt.Println(settings.Username)

	opts := mqtt.NewClientOptions()
	opts.AddBroker(settings.Broker)
	opts.SetClientID(settings.Id)
	opts.SetUsername(settings.Username)
	opts.SetPassword(password)

	opts.SetCleanSession(settings.CleanSession)

	if settings.Store != "" && settings.Store != ":memory:" {
		logger.Debugf("Using file store: %s", settings.Store)
		opts.SetStore(mqtt.NewFileStore(settings.Store))
	}
	fmt.Println(opts)

	return opts
}
