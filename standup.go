package main

import (
	"errors"
	"github.com/nlopes/slack"
	"github.com/robfig/cron"
	"log"
)

type Standup struct {
	Name    string
	Cron    string
	Channel string
	Users   []string
	Rtm     *slack.RTM
}

var standupCron = cron.New()

//adds a new standup to the schedule
func addStandup(rtm *slack.RTM, args []string) error {
	if len(args) < 4 {
		return errors.New("Incorrect arguments. type 'help' for assistance")
	}
	standup := Standup{
		Name:    args[1],
		Cron:    args[2],
		Channel: args[3],
		Users:   args[4:],
		Rtm:     rtm,
	}
	log.Printf("Adding standup: %v", standup)
	standupCron.AddFunc("* * * * * *", standup.startStandup)
	log.Printf("%v", standupCron.Entries())
	return nil
}

//Sends standup to all defined users
func (s Standup) startStandup() {
	for _, user := range s.Users {
		_, _, channelID, err := api.OpenIMChannel(user)
		if err != nil {
			log.Println(err)
		}
		log.Printf("Sending standup to %s", channelID)
		s.Rtm.SendMessage(s.Rtm.NewOutgoingMessage("Test cron", channelID))
	}
}
