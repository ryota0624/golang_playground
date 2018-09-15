package main

import (
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/ryota0624/helloworld_log/common"
	"github.com/urfave/cli"

	"github.com/ryota0624/helloworld_log/sql_builder"
	_ "github.com/ryota0624/helloworld_log/statik"
	// _ "./statik"
)

func makeCLIApp() *cli.App {
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

func timeControll() {
	_ := time.Now()
	_, _ := time.Parse("2006-01-02", "2014-12-31")
	// https://qiita.com/taizo/items/acbee530bd33c803dab4
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
	cliApp := makeCLIApp()
	cliApp.Action = sql_builder.SQLBuilder{
		Output: os.Stdout,
	}.Run
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
