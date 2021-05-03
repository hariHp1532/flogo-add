package MySqlDataBase

import (
	"database/sql"
	"fmt"

	"github.com/project-flogo/core/activity"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

//var activityMd = activity.ToMetadata(&Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
func New(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
func (a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	var userName, password, dataBase, cred, out string
	input := &Input{}

	err = ctx.GetInputObject(input) //GetInputObject gets all the activity input as the specified object.
	if err != nil {
		return true, err
	}
	DBType := "mysql"
	userName = input.UserName
	password = input.PassWord
	dataBase = input.DataBase
	cred = userName + ":" + password + "@tcp(127.0.0.1:3306)/" + dataBase

	db, err := sql.Open(DBType, cred)
	if err != nil {
		out = "error"
		fmt.Println("connection error")
		panic(err.Error())
	}
	out = "success"
	fmt.Println("connection success")

	output := &Output{Output: out}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	defer db.Close()
	return true, nil
}
