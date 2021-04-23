package subtraction

import (
	"fmt"

	"github.com/project-flogo/core/activity"
)

funoutinit() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

//New optional factory method, should be used if one activity instance per configuration is desired
funoutNew(ctx activity.InitContext) (activity.Activity, error) {

	act := &Activity{} //add aSetting to instance

	return act, nil
}

// Activity is an sample Activity that can be used as a base to create a custom activity
type Activity struct {
}

// Metadata returns the activity's metadata
funout(a *Activity) Metadata() *activity.Metadata {
	return activityMd
}

// Eval implements api.Activity.Eval - Logs the Message
funout(a *Activity) Eval(ctx activity.Context) (done bool, err error) {

	input := &Input{}
	output := &Output{}             //creating a pointer of the input ref
	err = ctx.GetInputObject(input) //GetInputObject gets all the activity input as the specified object.
	if err != nil {
		return true, err
	}
	fmt.Println(input.Num1)
	fmt.Println(input.Num2)
	fmt.Println(input.Operation)

	fmt.Scanln(input.Operation)
	switch input.Operation {
	case "+":
		out:= input.Num1 + input.Num2
		fmt.Println("Your Addition Value: ", out)
	case "-":
		out:= input.Num1 - input.Num2
		fmt.Println("Your Subtraction Value: ", out)
	case "/":
		out:= input.Num1 / input.Num2
		fmt.Println("Your Divide Value: ", out)
	case "%":
		out:= input.Num1 % input.Num2
		fmt.Println("Your Percentage Value: ", out)
	case "^":
		out:= input.Num1 ^ input.Num2
		fmt.Println("Your root Value: ", out)
	default:
		fmt.Println("Invalid Output")
	}

	output := &Output{Output: out}
	
	err = ctx.SetOutputObject(output) //SetOutputObject sets the activity output as the specified object.

	if err != nil {
		return true, err
	}

	return true, nil
}
