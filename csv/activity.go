package csv

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

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

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	var start, out string

	input := &Input{}               //creating a pointer of the input ref
	err = ctx.GetInputObject(input) //GetInputObject gets all the activity input as the specified object.
	if err != nil {
		return true, err
	}
	records := readCsvFile("poc.csv")
	start = input.In
	start = records

	out = start

	output := &Output{Output: out}
	err = ctx.SetOutputObject(output) //SetOutputObject sets the activity output as the specified object.

	if err != nil {
		return true, err
	}

	return true, nil
}
