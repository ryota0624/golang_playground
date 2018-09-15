package sql_builder

import (
	"bytes"
	"io"
	"log"

	"github.com/golang/go/src/html/template"
	"github.com/urfave/cli"

	"github.com/ryota0624/helloworld_log/common"
	_ "github.com/ryota0624/helloworld_log/statik"

	// _ "./statik"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

const (
	WhereFlag   = "where"
	ColumnsFlag = "columns"
)

type SQLStruct struct {
	Columns string
	Table   string
	Where   string
}

func newSQLStruct(columns string, table string, where string) SQLStruct {
	return SQLStruct{
		Columns: columns,
		Table:   table,
		Where:   where,
	}
}

type SQLBuilder struct {
	Output io.Writer
}

func (builder SQLBuilder) Run(ctx *cli.Context) error {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	query, error := common.LoadStatikFS("/query.sql")

	if error != nil {
		panic(error)
	}

	sugar.Infof("query template %s", query)

	queryTemplate, err := template.New("sqlQuery").Parse(query)

	if err != nil {
		panic(err)
	}

	var doc bytes.Buffer

	tableName, ok := env["TABLE_NAME"]

	if !ok {
		panic("env[TABLE_NAME] is no defined")
	}

	data := newSQLStruct(ctx.String(ColumnsFlag), tableName, ctx.String(WhereFlag))

	if err := queryTemplate.Execute(&doc, data); err != nil {
		panic(err)
	}

	builder.Output.Write(doc.Bytes())

	return nil
}
