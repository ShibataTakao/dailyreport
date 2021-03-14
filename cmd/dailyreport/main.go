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
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "path",
				Usage:    "Path to daily report directory",
				Required: true,
				EnvVars:  []string{"DAILYREPORT_PATH"},
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "create",
				Usage: "Create today's daily report",
				Action: func(c *cli.Context) error {
					return dailyreport.Create(c.String("path"))
				},
			},
			{
				Name:  "validate",
				Usage: "Validate worktime in today's daily report",
				Flags: []cli.Flag{
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
						EnvVars:  []string{"DAILYREPORT_REPORT_START_DATE"},
					},
					&cli.StringFlag{
						Name:     "end-date",
						Aliases:  []string{"e"},
						Usage:    "End of date range. Format must be yyyymmdd.",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_REPORT_END_DATE"},
					},
					&cli.StringFlag{
						Name:     "backlog-api-key",
						Usage:    "Backlog api key",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_BACKLOG_API_KEY"},
					},
					&cli.StringFlag{
						Name:     "backlog-base-url",
						Usage:    "Backlog base URL",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_BACKLOG_BASE_URL"},
					},
					&cli.StringFlag{
						Name:     "queries",
						Aliases:  []string{"q"},
						Usage:    "Json array format queries to fetch backlog issues",
						Required: true,
						EnvVars:  []string{"DAILYREPORT_BACKLOG_QUERIES"},
					},
				},
				Action: func(c *cli.Context) error {
					return dailyreport.Report(c.String("path"), c.String("category"), c.String("start-date"), c.String("end-date"), c.String("backlog-api-key"), c.String("backlog-base-url"), c.String("queries"))
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
