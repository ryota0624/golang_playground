package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/go/src/html/template"
	"github.com/urfave/cli"

	_ "github.com/ryota0624/helloworld_log/statik"
	// _ "./statik"
	"github.com/joho/godotenv"
	"github.com/rakyll/statik/fs"
	"go.uber.org/zap"
)

const (
	whereFlag   = "where"
	columnsFlag = "columns"
)

func loadStatikFS(path string) (string, error) {

	statikFS, err := fs.New()

	if err != nil {
		return "", err
	}

	file, err := statikFS.Open(path)

	if err != nil {
		return "", err
	}

	queryAsByte, err := ioutil.ReadAll(file)

	if err != nil {
		return "", err
	}

	return string(queryAsByte), nil
}

func loggerSample() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	url := "http://localhost"
	sugar.Infof("Failed to fetch URL: %s", url)

	logic(*logger)
}

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

func makeCLIApp() *cli.App {
	cliApp := cli.NewApp()

	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  whereFlag,
			Value: "",
			Usage: "sql where",
		},
		cli.StringFlag{
			Name:  columnsFlag,
			Value: "*",
			Usage: "columns where",
		},
	}

	return cliApp
}

type SQLBuilter struct {
	output io.Writer
}

func (builder SQLBuilter) build(ctx *cli.Context) error {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	query, error := loadStatikFS("/query.sql")

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

	data := newSQLStruct(ctx.String(columnsFlag), tableName, ctx.String(whereFlag))

	if err := queryTemplate.Execute(&doc, data); err != nil {
		panic(err)
	}

	builder.output.Write(doc.Bytes())

	return nil
}

func main() {
	cliApp := makeCLIApp()
	cliApp.Action = SQLBuilter{
		output: os.Stdout,
	}.build
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type Record struct {
	Name string
}

func newRecord(name string) Record {
	record := Record{}
	record.Name = name
	return record
}

func logic(logger zap.Logger) {
	logger.Info("startLogic")

	array := "HOUGE" + "HUGE"

	logger.Info(array)
	logger.Info("record")

	logger.Info("endLogic")
}
