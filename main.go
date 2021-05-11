package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	DBType := "mysql"
	userName := "root"
	password := "admin"
	dataBase := "csv"
	cred := userName + ":" + password + "@tcp(127.0.0.1:3306)/" + dataBase
	db, err := sql.Open(DBType, cred)

	if err != nil {
		log.Fatal(err)
	}

	//INSERT INTO mqtt (id,message)

	//INSERT INTO mqtt(id,message) VALUES (1,'hari')
	val := "hari"
	sql := `INSERT INTO mqtt VALUES (1,` + val + `);`

	//sql := a + val + ")"
	res, errr := db.Exec(sql)

	if errr != nil {
		panic(errr.Error())
	}

	lastId, errr := res.LastInsertId()

	if errr != nil {
		log.Fatal(errr)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}
