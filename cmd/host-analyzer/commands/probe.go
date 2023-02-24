package commands

import (
	"fmt"

	"github.com/pimg/host-analyzer/internal/host"
	"github.com/urfave/cli/v2"
)

func init() {
	fmt.Println("Initializing probe commands")
	GetRegistry().RegisterCommands(ProbeCommands())
}

func ProbeCommands() []*cli.Command {
	probeFlags := GetRegistry().GlobalFlags //when additional flags are needed these can be appended here
	probeCommands := []*cli.Command{
		{
			Name:  "icmp",
			Usage: "Performs an ICMP (ping) on a particual Host",
			Flags: probeFlags,
			Action: func(c *cli.Context) error {
				return host.PingHost(c.String("host"))
			},
		},
		{
			Name:  "http",
			Usage: "Performs an HTTP call and finds HTTP protocol information",
			Flags: probeFlags,
			Action: func(c *cli.Context) error {
				return host.ProbeHTTP(c.String("host"))
			},
		},
		{
			Name:  "tls",
			Usage: "Performs a TLS handshake to find information about the TLS configuration",
			Flags: probeFlags,
			Action: func(c *cli.Context) error {
				return host.ProbeTLS(c.String("host"), c.Int("port"))
			},
		},
	}

	return probeCommands
}
