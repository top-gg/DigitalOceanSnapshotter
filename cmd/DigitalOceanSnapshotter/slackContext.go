package main

import "github.com/slack-go/slack"

// SlackContext is an helper struct to acess slack actions
type SlackContext struct {
	client    *slack.Client
	channelID string
}

// SendMessageWithContent sends a message with content to the pre defined channel
func (s SlackContext) SendMessageWithContent(content string) error {
	_, _, _, err := s.client.SendMessage(s.channelID, slack.MsgOptionText(content, true))
	return err
}

// SendMessageWithEmbed sends a message with a Rich Embed to the pre defined channel
func (s SlackContext) SendMessageWithEmbed(attachment slack.Attachment) error {
	_, _, _, err := s.client.SendMessage(s.channelID, slack.MsgOptionAttachments(attachment))
	return err
}
