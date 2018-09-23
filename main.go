package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/ryota0624/helloworld_log/bqsample"
	"github.com/ryota0624/helloworld_log/common"
	"github.com/ryota0624/helloworld_log/record_sample"
	"github.com/ryota0624/helloworld_log/record_sample/nest"
	"github.com/ryota0624/helloworld_log/sql_builder"
	_ "github.com/ryota0624/helloworld_log/statik"
	"github.com/urfave/cli"
	// _ "./statik"
)

func makeCLIApp() *cli.App {
	log.Printf("CONFIG_ENV: %s", os.Getenv("CONFIG_ENV"))

	cliApp := cli.NewApp()

	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  sql_builder.WhereFlag,
			Value: "",
			Usage: "sql where",
		},
		cli.StringFlag{
			Name:  sql_builder.ColumnsFlag,
			Value: "*",
			Usage: "columns where",
		},
	}

	return cliApp
}

func nowString() string {
	now := time.Now()
	_, _ = time.Parse("2006-01-02", "2014-12-31")
	// https://qiita.com/taizo/items/acbee530bd33c803dab4
	return now.Format("20060102")
}

func yeasterdayString() string {
	yesterday := time.Now().AddDate(0, 0, -1)
	return yesterday.Format("20060102")
}

func execShell() {
	configShell, loadConfigError := common.LoadStatikFS("/config.sh")
	if loadConfigError != nil {
		panic(loadConfigError)
	}

	_, execError := exec.Command("sh", "-c", configShell).Output()
	println(os.Getenv("CONFIG2"))
	if execError != nil {
		panic(execError)
	}
}

func main() {
	bqsample.Do()

	os.Exit(0)
	nest.Fn()
	_ = record_sample.Record{}
	log.Println("now", nowString())
	log.Println("yesterday", yeasterdayString())

	cliApp := makeCLIApp()
	cliApp.Action = sql_builder.SQLBuilder{
		Output: os.Stdout,
	}.Run
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
