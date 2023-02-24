package commands

import (
	"github.com/pimg/host-analyzer/internal/host"
	"github.com/urfave/cli/v2"
)

func ProbeCommands(commands *[]*cli.Command, hostAnalyzerFlags []cli.Flag) {
	probeCommands := []*cli.Command{
		{
			Name:  "icmp",
			Usage: "Performs an ICMP (ping) on a particual Host",
			Flags: hostAnalyzerFlags,
			Action: func(c *cli.Context) error {
				return host.PingHost(c.String("host"))
			},
		},
		{
			Name:  "http",
			Usage: "Performs an HTTP call and finds HTTP protocol information",
			Flags: hostAnalyzerFlags,
			Action: func(c *cli.Context) error {
				return host.ProbeHTTP(c.String("host"))
			},
		},
		{
			Name:  "tls",
			Usage: "Performs a TLS handshake to find information about the TLS configuration",
			Flags: hostAnalyzerFlags,
			Action: func(c *cli.Context) error {
				return host.ProbeTLS(c.String("host"), c.Int("port"))
			},
		},
	}

	*commands = append(*commands, probeCommands...)
}
