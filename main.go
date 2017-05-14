package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"
	"strings"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	if token == "" {
		log.Println("SLACK_TOKEN not found")
		os.Exit(1)
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"))
	//api.SetDebug(true)
	log.Println("Slack Bot Starting")

	rtm := api.NewRTM()
	go rtm.ManageConnection()

	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {
		case *slack.MessageEvent:
			info := rtm.GetInfo()
			if ev.User != info.User.ID { //verifies the message is not from the bot
				log.Printf("%v", ev.Text)
				if verifyChannelisIM(rtm, ev.Channel) {
					log.Printf("Direct message: %s", ev.Text)
					go respond(rtm, ev)
				}
			}
		case *slack.RTMError:
			log.Printf("Error: %s\n", ev.Error())
		case *slack.InvalidAuthEvent:
			log.Printf("Invalid credentials")
			os.Exit(1)
		default:
			// ignore other events
		}
	}
}

//function to handle direct messages to bot
func respond(rtm *slack.RTM, msg *slack.MessageEvent) {
	var response string
	trimmedCmd := strings.TrimSpace(msg.Text)
	args := strings.Split(trimmedCmd, " ")
	switch args[0] {
	case "help":
		helpstr := help()
		response = fmt.Sprintf(helpstr)
	case "add_standup":
		response = args[0] + ": Command not implemented"
	case "add_user":
		response = args[0] + ": Command not implemented"
	case "standup_info":
		response = args[0] + ": Command not implemented"
	default:
		helpstr := args[0] + ": Command not found\n"
		helpstr += help()
		response = fmt.Sprintf(helpstr)
	}
	rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
}

func help() string {
	helpstr := "Supported Commands:\n"
	helpstr += "> add_standup <standup name> <cron expression> <channel> <user1> <userN>\n"
	helpstr += "> add_user <standup name> <user>\n"
	helpstr += "> standup_info <standup name>\n"
	return helpstr
}

func verifyChannelisIM(rtm *slack.RTM, id string) bool {
	info := rtm.GetInfo()
	for _, im := range info.IMs {
		if im.ID == id {
			return true
		}
	}
	return false
}
