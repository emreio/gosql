package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/emreio/golib"
	"github.com/emreio/sqlgo/types"
	_ "github.com/microsoft/go-mssqldb"
)

const sql_ado string = "server={SERVER};user id=${USER};password={PASS};database=Experimantal;app name=Go App"

func CallAddInterface(mUtil golib.Multiply) {
	mUtil.Multiply(12, 15)
}

var db *sql.DB

func main() {
	db, _err := sql.Open("sqlserver", sql_ado)
	if _err != nil {
		fmt.Println(_err)
	}

	context := context.Background()

	rows, _err := db.QueryContext(context, "SELECT * FROM dbo.Cities")
	if _err != nil {
		log.Fatal(_err.Error())
	}

	var ross string

	for rows.Next() {
		rows.Scan(&ross)
		fmt.Println(ross)
	}
	ii := &golib.MathUtil{}
	ii.LoggingEnabled = true
	CallAddInterface(ii)
	CallSql(&types.SqlRepository{})
}

func CallSql(repo types.Repository) {
	repo.Select("SELECT * FROM dbo.* **")
}
