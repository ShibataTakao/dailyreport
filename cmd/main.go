package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ShibataTakao/dailyreport/dailyreport"
	"github.com/spf13/cobra"
)

func newCommand() *cobra.Command {
	var (
		dir      string
		startStr string
		endStr   string
	)

	command := &cobra.Command{
		Use:   "dailyreport",
		Short: "Aggregate daily reports.",
		RunE: func(_ *cobra.Command, _ []string) error {
			repository := dailyreport.NewRepository(dir)
			app := dailyreport.NewApplicationService(repository)
			start, err := time.Parse("20060102", startStr)
			if err != nil {
				return err
			}
			end, err := time.Parse("20060102", endStr)
			if err != nil {
				return err
			}
			aggregated, err := app.Read(start, end)
			if err != nil {
				return err
			}
			fmt.Println(aggregated)
			return nil
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	now := time.Now()
	command.Flags().StringVar(&dir, "dir", ".", "The directory where daily report files are stored.")
	command.Flags().StringVarP(&startStr, "start", "s", now.Format("20060102"), "The start date of the daily report range.")
	command.Flags().StringVarP(&endStr, "end", "e", now.Add(-720*time.Hour).Format("20060102"), "The end date of the daily report range.")

	return command
}

func main() {
	log.SetFlags(0)
	if err := newCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
