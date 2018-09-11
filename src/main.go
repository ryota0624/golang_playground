package main

import "go.uber.org/zap"

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	url := "http://localhost"
	sugar.Infof("Failed to fetch URL: %s", url)
}
