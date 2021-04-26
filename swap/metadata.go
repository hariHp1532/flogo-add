package swap

import (
	"github.com/project-flogo/core/data/coerce"
	//"github.com/spf13/cast"
)

type Input struct {
	Number1 int `md:"Number1.required"`
	Number2 int `md:"Number2.required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	Val1, _ := coerce.ToInt(values["Number1"])
	r.Number1 = Val1

	Val2, _ := coerce.ToInt(values["Number2"])
	r.Number2 = Val2
	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Number1": r.Number1,
		"Number2": r.Number2,
	}
}

type Output struct {
	Output1 int `md:"Output1"`
	Output2 int `md:"Output2"`
}

func (o *Output) FromMap(values map[string]interface{}) error {
	strVal, _ := coerce.ToInt(values["Output1"])
	o.Output1 = strVal

	strVa2, _ := coerce.ToInt(values["Output2"])
	o.Output2 = strVa2
	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output1": o.Output1,
		"Output2": o.Output2,
	}
}
