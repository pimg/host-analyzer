package main

import (
	"log"
	"os"

	commands "github.com/pimg/host-analyzer/cmd/host-analyzer/commands"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "hostname analyzer cli"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name servers"

	hostAnalyzerFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "",
		},
		&cli.IntFlag{
			Name:  "port",
			Value: 443,
		},
	}

	commands.HostCommands(&app.Commands, hostAnalyzerFlags)
	commands.ProbeCommands(&app.Commands, hostAnalyzerFlags)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
