package cmd

import (
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/Dacode45/go-timekeeper/timekeeping"
	"github.com/urfave/cli"
)

var StartTaskCommand = cli.Command{
	Name:     "start",
	Usage:    "Starts a Task",
	Flags:    VerboseFlags,
	Action:   StartTask,
	Category: "Basic Operations",
}

// StartTask starts a task
func StartTask(c *cli.Context) error {
	name := c.Args().First()
	if name == "" {
		return fmt.Errorf("Name is required")
	}
	start := time.Now()
	timekeeping.NewTask(name, "", start)
	log.Infof("Starting %q: %s", name, start.Format(TimeFormat))
	return nil
}
