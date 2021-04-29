package logic

import (
	"fmt"

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

// Eval implements api.Activity.Eval - Logs the Message
func (a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{} //creating a pointer of the input ref
	var out int
	err = ctx.GetInputObject(input) //GetInputObject gets all the activity input as the specified object.
	if err != nil {
		return true, err
	}
	fmt.Println("num1 value:", input.Num1)
	fmt.Println("num2 value:", input.Num2)
	fmt.Println("operation:", input.Operation)

	fmt.Scanln(input.Operation)

	switch input.Operation {
	case "AND":
		out = input.Num1 & input.Num2
		fmt.Println("Result of num1 & num2 = %d", out)
	case "OR":
		out = input.Num1 | input.Num2
		fmt.Println("Result of num1 | num2 = %d", out)
	case "XOR":
		out = input.Num1 ^ input.Num2
		fmt.Println("Result of num1 ^ num2 = %d", out)
	case "left-shift":
		out = input.Num1 << 1
		fmt.Println("Result of num1 << 1 = %d", out)
	case "right-shift":
		out = input.Num1 >> 1
		fmt.Println("Result of num1 >> 1 = %d", out)
	case "NOT":
		out = input.Num1 &^ input.Num2
		fmt.Println("Result of num1 &^ num2 = %d", out)
	default:
		fmt.Println("Invalid Operation")
	}

	output := &Output{Output: out}
	err = ctx.SetOutputObject(output) //SetOutputObject sets the activity output as the specified object.

	if err != nil {
		return true, err
	}

	return true, nil
}
