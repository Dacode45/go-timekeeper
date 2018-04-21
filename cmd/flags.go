package cmd

import "github.com/urfave/cli"

// VerboseFlags allow the user to set the verbosity
var VerboseFlags = []cli.Flag{
	cli.BoolFlag{
		Name: "verbose, v",
	},
}

var ClosedFlag = cli.BoolFlag{
	Name:  "closed",
	Usage: "Only closed tasks",
}

var OpenFlag = cli.BoolFlag{
	Name:  "open",
	Usage: "Only open tasks",
}
