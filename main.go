package main

import (
	"log"

	"github.com/ShibataTakao/dailyreport/cmd"
)

func main() {
	if err := cmd.NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
