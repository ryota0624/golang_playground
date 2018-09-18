package recordsample

import "go.uber.org/zap"

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

func loggerSample() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()

	url := "http://localhost"
	sugar.Infof("Failed to fetch URL: %s", url)

	logic(*logger)
}
