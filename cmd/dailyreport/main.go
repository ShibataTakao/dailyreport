package main

import (
	"log"
	"os"

	"github.com/shibataka000/daily-report/internal/dailyreport"
	"github.com/urfave/cli/v2"
)

func main() {
	log.SetFlags(0)

	app := &cli.App{
		Name:  "dailyreport",
		Usage: "Dailyreport manager",
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create today's daily report",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "template",
						Usage:    "Path to daily report template file",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "dir",
						Usage:    "Path to daily report directory",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Create(c.String("template"), c.String("dir"))
				},
			},
			{
				Name:  "validate",
				Usage: "Validate worktime in today's daily report",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "dir",
						Usage:    "Path to daily report directory",
						Required: true,
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Validate(c.String("dir"))
				},
			},
			{
				Name:  "kkms",
				Usage: "",
				Action: func(c *cli.Context) error {
					return dailyreport.KKMS()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
