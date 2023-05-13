package main

import (
	"log"

	"github.com/ShibataTakao/worklog/cmd"
)

func main() {
	if err := cmd.NewCommand().Execute(); err != nil {
		log.Fatal(err)
	}
}
