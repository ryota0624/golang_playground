package mysqlsample

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Do() {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}
	db.Exec("load file.csv.gzip")
}
