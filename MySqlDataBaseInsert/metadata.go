package MySqlDataBaseInsert

import (
	"github.com/project-flogo/core/data/coerce"
	//"github.com/spf13/cast"
)

type Input struct {
	UserName string `md:"UserName.required"`
	PassWord string `md:"PassWord.required"`
	DataBase string `md:"DataBase.required"`
	Table    string `md:"Table.required"`
	Value    string `md:"Value.required"`
}

func (r *Input) FromMap(values map[string]interface{}) error {

	Val1, _ := coerce.ToString(values["UserName"])
	r.UserName = Val1

	Val2, _ := coerce.ToString(values["PassWord"])
	r.PassWord = Val2

	Val3, _ := coerce.ToString(values["DataBase"])
	r.DataBase = Val3

	Val4, _ := coerce.ToString(values["Table"])
	r.Table = Val4

	Val5, _ := coerce.ToString(values["Value"])
	r.Value = Val5

	return nil
}

func (r *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"UserName": r.UserName,
		"PassWord": r.PassWord,
		"DataBase": r.DataBase,
		"Table":    r.Table,
		"Value":    r.Value,
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
