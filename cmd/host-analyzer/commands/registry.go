package commands

import "github.com/urfave/cli/v2"

var commandRegistry = &Registry{
	ImplementedCommands: []*cli.Command{},
	GlobalFlags: []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "",
		},
	},
}

type Registry struct {
	ImplementedCommands []*cli.Command
	GlobalFlags         []cli.Flag
}

func GetRegistry() *Registry {
	return commandRegistry
}

func (registry *Registry) RegisterCommands(commands []*cli.Command) {
	registry.ImplementedCommands = append(commandRegistry.ImplementedCommands, commands...)
}
