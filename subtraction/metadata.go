package subtraction

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
	//"github.com/spf13/cast"
)

type Input struct {
	Num1 int `md:"Num1.required"`
	Num2 int `md:"Num2.required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	Val1, _ := coerce.ToInt(values["Num1"])
	r.Num1 = Val1
	fmt.Println(values["Num1"])
	Val2, _ := coerce.ToInt(values["Num2"])
	r.Num2 = Val2
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Num1": r.Num1,
		"Num2": r.Num2,
	}
}

type Output struct {
	Output int `md:"Output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToInt(values["Output"])
	o.Output = strVal
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
	}
} 
