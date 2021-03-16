package main

import (
	"log"
	"os"

	"github.com/gba-3/goslack"
)

func main() {
	webhookURL := os.Getenv("SLACK_URL")
	cli := goslack.NewClient(webhookURL)
	err := cli.Post("testmsg")
	if err != nil {
		log.Println(err.Error())
	}
}
