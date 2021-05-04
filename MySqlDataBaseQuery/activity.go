package MySqlDataBaseQuery

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/project-flogo/core/activity"
)

func init() {
	_ = activity.Register(&Activity{}) //activity.Register(&Activity{}, New) to create instances using factory method 'New'
}

var activityMd = activity.ToMetadata(&Input{}, &Output{})

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

	var DBType, userName, password, dataBase, cred, querys, typequery, out string
	input := &Input{}

	err = ctx.GetInputObject(input) //GetInputObject gets all the activity input as the specified object.
	if err != nil {
		return true, err
	}
	DBType = "mysql"
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
	//out = "success"
	fmt.Println("connection success")

	querys = input.Querys
	typequery = input.TypeQuery

	fmt.Print("Select The Query statement insert, delete: \n")
	fmt.Scanln(&querys)

	switch querys {
	case "insert":
		Qsql := typequery
		res, errr := db.Exec(Qsql)
		if errr != nil {
			panic(errr.Error())
		}
		out, errr := res.LastInsertId()
		if errr != nil {
			log.Fatal(errr)
		}
		fmt.Printf("The last inserted row id: %d\n", out)

	case "delete":
		qsql := typequery
		res, err := db.Exec(qsql)
		if err != nil {
			panic(err.Error())
		}
		out, err := res.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("The statement affected %d rows\n", out)

	default:
		fmt.Println("Invalid Query select")
	}

	output := &Output{Output: out}

	err = ctx.SetOutputObject(output)
	if err != nil {
		return true, err
	}

	defer db.Close()
	return true, nil
}
