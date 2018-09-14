package main

import (
	"bytes"
	"io/ioutil"
	"text/template"

	_ "./statik"

	"github.com/rakyll/statik/fs"
	"go.uber.org/zap"
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
}

func newSqlStruct(columns string, table string) SQLStruct {
	return SQLStruct{
		Columns: columns,
		Table:   table,
	}
}

func main() {
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

	data := newSqlStruct("id,name", "users")

	if err := queryTemplate.Execute(&doc, data); err != nil {
		panic(err)
	}

	s := doc.String()

	sugar.Infof("destination query %s", s)
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
	// record := newRecord("new!")
	logger.Info("record")

	logger.Info("endLogic")
}
