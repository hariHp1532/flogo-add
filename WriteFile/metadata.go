package WriteFile

import (
	"github.com/project-flogo/core/data/coerce"
	//"github.com/spf13/cast"
)

type Input struct {
	FileName string `md:"FileName.required"`
	Data     string `md:"Data.required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	Val1, _ := coerce.ToString(values["FileName"])
	r.FileName = Val1

	Val2, _ := coerce.ToString(values["Data"])
	r.Data = Val2

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"FileName": r.FileName,
		"Data":     r.Data,
	}
}

type Output struct {
	Output string `md:"Output"`
}

func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	o.Output, err = coerce.ToString(values["Output"])
	if err != nil {
		return err
	}

	return nil
}

func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Output": o.Output,
	}
}
