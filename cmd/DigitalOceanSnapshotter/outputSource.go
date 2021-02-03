package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/slack-go/slack"
)

// OutputSource is an abstraction for outputting specific events to other services (e.g. Discord, Slack or Whatsapp)
type OutputSource interface {
	SendEvent(string, log.Level) error
}

// SendEvent forwards event to Slack
func (s SlackContext) SendEvent(content string, level log.Level) error {

	color := "#00FF00"

	if level == log.ErrorLevel {
		color = "#FF0000"
	}

	return s.SendMessageWithEmbed(slack.Attachment{
		Color:      color,
		AuthorName: "DigitalOceanSnapshotter",
		AuthorIcon: "https://cdn.top.gg/icons/DO_Logo_icon_blue.png",
		Text:       content,
		Title:      "DigitalOceanSnapshotter",
		TitleLink:  "https://github.com/top-gg/DigitalOceanSnapshotter",
	})
}
