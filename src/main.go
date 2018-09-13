package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "http://localhost"
	sugar.Infof("Failed to fetch URL: %s", url)

	logic(*logger)
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
