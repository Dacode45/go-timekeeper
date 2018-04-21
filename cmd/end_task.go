package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/Dacode45/go-timekeeper/timekeeping"
	log "github.com/sirupsen/logrus"
	input "github.com/tcnksm/go-input"
	"github.com/urfave/cli"
)

var EndTaskCommand = cli.Command{
	Name:     "end",
	Usage:    "Ends a tasks",
	Flags:    VerboseFlags,
	Action:   EndTask,
	Category: "Basic Operations",
}

// EndTask ends a task
func EndTask(c *cli.Context) error {
	name := c.Args().First()
	if name == "" {
		return fmt.Errorf("Name is required")
	}
	end := time.Now()
	candidates := timekeeping.FindTask(name)
	open := make([]timekeeping.Task, 0, len(candidates))
	for _, c := range candidates {
		if !c.IsClosed {
			open = append(open, c)
		}
	}
	if len(open) == 0 {
		return fmt.Errorf("No open task with name %q", name)
	} else if len(open) == 1 {
		task := open[0]
		task.EndTask(end)
		log.Infof("Ending %q: %s", name, end.Format(TimeFormat))
		return nil
	}

	options := make([]string, len(open))
	for _, c := range open {
		options = append(options, fmt.Sprintf("%q: %s", c.Name, c.Intervals[len(c.Intervals)-1].Start.Format(TimeFormat)))
	}

	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}
	confirm, err := ui.Select(
		"Which task would you like to close",
		options,
		&input.Options{
			Default: options[0],
			Loop:    true,
		},
	)
	if err != nil {
		return err
	}
	var selected int
	for i := range options {
		if options[i] == confirm {
			selected = i
			break
		}
	}

	task := open[selected]
	task.EndTask(end)
	log.Infof("Ending %q: %s", name, end.Format(TimeFormat))
	return nil
}
