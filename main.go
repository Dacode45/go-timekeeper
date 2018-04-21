package main

import (
	"log"
	"os"

	"github.com/Dacode45/go-timekeeper/cmd"
	"github.com/Dacode45/go-timekeeper/timekeeping"
	_ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli"
)

func initApp() *cli.App {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		cmd.StartTaskCommand,
		cmd.EndTaskCommand,
		cmd.ListTasksCommand,
	}
	return app
}

func main() {
	if err := timekeeping.Setup(); err != nil {
		log.Fatal(err)
	}
	defer timekeeping.Cleanup()

	app := initApp()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
