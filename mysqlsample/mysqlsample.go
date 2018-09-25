package mysqlsample

import (
	"context"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

func Do(ctx context.Context) {
	db, err := sql.Open("mysql", "user:password@/dbname")
	if err != nil {
		panic(err)
	}

	tx, beginTxError := db.BeginTx(ctx, nil) // transactionの生成

	if beginTxError != nil {
		panic(beginTxError)
	}

	mysql.RegisterLocalFile("csv.gzip")                   // load するファイルを許容するようにする
	_, execLoadFileError := tx.Exec("load file.csv.gzip") // loadの実行
	if execLoadFileError != nil {
		panic(execLoadFileError)
	}

	txCommitError := tx.Commit()
	if txCommitError != nil {
		panic(txCommitError)
	}
}
