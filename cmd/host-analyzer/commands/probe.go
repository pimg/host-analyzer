package commands

import (
	"github.com/pimg/host-analyzer/internal/host"
	"github.com/urfave/cli/v2"
)

func init() {
	GetRegistry().RegisterCommands(ProbeCommands())
}

func ProbeCommands() []*cli.Command {
	probeFlags := GetRegistry().GlobalFlags //when additional flags are needed these can be appended here

	schemeFlag := &cli.StringFlag{
		Name:  "scheme",
		Value: "https",
	}

	portFlag := &cli.StringFlag{
		Name:  "port",
		Value: "443",
	}

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
			Flags: append(probeFlags, schemeFlag),
			Action: func(c *cli.Context) error {
				return host.ProbeHTTP(c.String("scheme"), c.String("host"), c.String("port"))
			},
		},
		{
			Name:  "tls",
			Usage: "Performs a TLS handshake to find information about the TLS configuration",
			Flags: append(probeFlags, portFlag),
			Action: func(c *cli.Context) error {
				return host.ProbeTLS(c.String("host"), c.Int("port"))
			},
		},
	}

	return probeCommands
}
