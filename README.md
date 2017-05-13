# Sitdown-Standup

A Slack bot for doing asynchronous text-based standups.

Future Features:
- Multiple standups with one bot
- Schedule based standup notifications, customizable per team
- Allow pausing of standup and use vacations
- Allow editing of completed standups from user
- Status reports on standups completion
- Open to any suggestions

This is very much a work in progress and just a fun side-project for now.

## Scheduling Idea

My intial idea is to use https://godoc.org/github.com/robfig/cron for scheduling the stand ups
The user messages the bot with "add $cron_expression $user1 $user2 $user3" 
This will add a standup that messages all users based on the cron_expression.


