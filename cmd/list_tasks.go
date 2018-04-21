package cmd

import (
	"html/template"
	"os"

	"github.com/Dacode45/go-timekeeper/timekeeping"
	"github.com/urfave/cli"
)

var ListTasksFlags = append(
	[]cli.Flag{},
	OpenFlag,
	ClosedFlag,
)

var ListTasksCommand = cli.Command{
	Name:     "list",
	Usage:    "list all tasks",
	Flags:    append(ListTasksFlags, VerboseFlags...),
	Action:   ListTasks,
	Category: "Inspection",
}

func ListTasks(c *cli.Context) error {
	filterOpen := c.Bool("open") || c.Bool("closed")
	t := template.New("list tasks")
	t, _ = t.Parse(listTemplate)
	var candidates []timekeeping.Task
	if !filterOpen {
		candidates = timekeeping.ListTask()
	} else {
		candidates = timekeeping.ListClosedTask(c.Bool("closed") && !c.Bool("open"))
	}
	t.Execute(os.Stdout, candidates)
	return nil
}

const listTemplate = `Tasks: 
{{range .}}
	{{.TaskID}}	{{.Name}}:{{if .IsClosed}} Closed {{else}} Open {{end}}
		{{range .Intervals}}
			Start: {{.Start.Format "Jan 02, 2006 15:04:05 UTC"}}
			{{if .IsClosed}}End: {{.End.Format "Jan 02, 2006 15:04:05 UTC"}}
		{{end}}
{{end}}
`
