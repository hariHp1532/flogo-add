package csv

import (
	"fmt"

	"github.com/project-flogo/core/data/coerce"
	//"github.com/spf13/cast"
)

type Input struct {
	In string `md:"In.required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	Val1, _ := coerce.ToString(values["In"])
	r.In = Val1
	fmt.Println("num1 value is: ", values["In"])

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"In": r.In,
	}
}

type Output struct {
	Output string `md:"Output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToString(values["Output"])
	o.Output = strVal
	//fmt.Println("output value is: ", values["Output"])
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
	}
}
