package main

import (
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
		helpResp(rtm, msg.Channel)
	case "add_standup":
		addStandup(args[1], args[2], args[3], args[4:])
		response = "Standup Added"
	case "add_user":
		response = args[0] + ": Command not implemented"
	case "standup_info":
		response = args[0] + ": Command not implemented"
	}
	if response != "" {
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}

//handles the help block
func helpResp(rtm *slack.RTM, channel string) {
	commands := map[string]string{
		"add_standup <name> <cron> <channel> <user1> <userN>": "Creates a standup that will message all given users based on cron schedule",
		"add_user <name> <user>":                              "Adds a user to the given standup",
		"list_standups":                                       "Lists all standups with schedule and users",
	}
	fields := make([]slack.AttachmentField, 0)
	for k, v := range commands {
		fields = append(fields, slack.AttachmentField{
			Title: k,
			Value: v,
		})
	}
	attachment := slack.Attachment{
		Pretext: "Supported Commands:",
		Color:   "#B733FF",
		Fields:  fields,
	}
	params := slack.PostMessageParameters{}
	params.AsUser = true
	params.Attachments = []slack.Attachment{attachment}
	rtm.PostMessage(channel, "", params)
}

//checks that a channel is an IM
func verifyChannelisIM(rtm *slack.RTM, id string) bool {
	info := rtm.GetInfo()
	for _, im := range info.IMs {
		if im.ID == id {
			return true
		}
	}
	return false
}
