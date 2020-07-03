package main

import (
	"log"
	"os"

	"github.com/pimg/host-analyzer/internal/host"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "hostname analyzer cli"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name servers"

	hostAnalyzerFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "",
		},
		cli.IntFlag{
			Name:  "port",
			Value: 443,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the Name Servers for a particular host",
			Flags: hostAnalyzerFlags,
			Action: func(c *cli.Context) error {
				return host.FindNameServers(c.String("host"))
			},
		},
		{
			Name:  "ip",
			Usage: "Looks up the IP addresses for a particular host",
			Flags: hostAnalyzerFlags,
			Action: func(c *cli.Context) error {
				return host.FindIPAddresses(c.String("host"))
			},
		},
		{
			Name:  "cname",
			Usage: "Looks up the CNAME for a particular Host",
			Flags: hostAnalyzerFlags,
			Action: func(c *cli.Context) error {
				return host.FindCNAME(c.String("host"))
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up the MX records for a particular Host",
			Flags: hostAnalyzerFlags,
			Action: func(c *cli.Context) error {
				return host.FindMXRecords(c.String("host"))
			},
		},
		{
			Name:  "txt",
			Usage: "Looks up the TXT records for a particular Host",
			Flags: hostAnalyzerFlags,
			Action: func(c *cli.Context) error {
				return host.FindTXTRecords(c.String("host"))
			},
		},
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
