package main

import (
	"log"
)

func addStandup(suName string, cronStr string, channel string, users []string) {
	log.Printf("Adding standup: %s %s %s %v", suName, cronStr, channel, users)
}
