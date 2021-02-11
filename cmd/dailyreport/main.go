package main

import (
	"log"
	"os"

	"github.com/ShibataTakao/daily-report/internal/dailyreport"
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
						Name:     "path",
						Usage:    "Path to daily report directory",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_PATH"},
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Create(c.String("path"))
				},
			},
			{
				Name:  "validate",
				Usage: "Validate worktime in today's daily report",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Path to daily report directory",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_PATH"},
					},
					&cli.StringFlag{
						Name:     "filepath",
						Aliases:  []string{"f"},
						Usage:    "Path to daily report file. This is conflict to 'path'",
						Required: false,
						Value:    "",
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Validate(c.String("path"), c.String("filepath"))
				},
			},
			{
				Name:  "report",
				Usage: "Report tasks in specific category from daily reports and issues",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "path",
						Usage:    "Path to daily report directory",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_PATH"},
					},
					&cli.StringFlag{
						Name:     "category",
						Usage:    "Task category to report",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_REPORT_CATEGORY"},
					},
					&cli.StringFlag{
						Name:     "start-date",
						Aliases:  []string{"s"},
						Usage:    "Beginning of date range. Format must be yyyymmdd.",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "end-date",
						Aliases:  []string{"e"},
						Usage:    "End of date range. Format must be yyyymmdd.",
						Required: true,
					},
					&cli.StringFlag{
						Name:     "trello-app-key",
						Usage:    "Trello app key",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_TRELLO_APP_KEY"},
					},
					&cli.StringFlag{
						Name:     "trello-token",
						Usage:    "Trello token",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_TRELLO_TOKEN"},
					},
					&cli.StringFlag{
						Name:     "queries",
						Aliases:  []string{"q"},
						Usage:    "Comma separated queries to fetch trello cards",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_TRELLO_QUERIES"},
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Report(c.String("path"), c.String("category"), c.String("start-date"), c.String("end-date"), c.String("trello-app-key"), c.String("trello-token"), c.String("queries"))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
