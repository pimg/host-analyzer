package commands

import (
	"fmt"

	"github.com/pimg/host-analyzer/internal/host"
	"github.com/urfave/cli/v2"
)

func init() {
	fmt.Println("Initializing host commands")
	GetRegistry().RegisterCommands(hostCommand())
}

func hostCommand() []*cli.Command {
	hostCommandFlags := GetRegistry().GlobalFlags //when additional flags are needed these can be appended here
	hostCommands := []*cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a particular host",
			Flags: hostCommandFlags,
			Action: func(c *cli.Context) error {
				return host.FindNameServers(c.String("host"))
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: hostCommandFlags,
			Action: func(c *cli.Context) error {
				return host.FindIPAddresses(c.String("host"))
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular Host",
			Flags: hostCommandFlags,
			Action: func(c *cli.Context) error {
				return host.FindCNAME(c.String("host"))
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular Host",
			Flags: hostCommandFlags,
			Action: func(c *cli.Context) error {
				return host.FindMXRecords(c.String("host"))
			},
		},
		{
			Name:  "txt",
			Usage: "Looks up the TXT records for a particular Host",
			Flags: hostCommandFlags,
			Action: func(c *cli.Context) error {
				return host.FindTXTRecords(c.String("host"))
			},
		},
	}
	return hostCommands
}
