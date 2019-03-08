package slack

import (
	"os"

	"github.com/golang/glog"
	"github.com/nlopes/slack"
)

func PostToSlack(text slack.MsgOption, attachment slack.Attachment) {
	channelID := os.Getenv("SLACK_CHANNEL")
	token := os.Getenv("SLACK_TOKEN")

	if channelID == "" || token == "" {
		glog.V(0).Infof("Please set environment variables SLACK_TOKEN and SLACK_CHANNEL to send a slack message.")
		return
	}

	api := slack.New(token)

	_, _, err := api.PostMessage(
		channelID,
		slack.MsgOptionUsername("Packt Bot"),
		slack.MsgOptionIconURL("http://krikkit.gration.org/images/bender_small.jpg"),
		text,
		slack.MsgOptionAttachments(attachment),
	)
	if err != nil {
		glog.V(0).Infof("Got error from Slack post: %s", err)
	}
}
