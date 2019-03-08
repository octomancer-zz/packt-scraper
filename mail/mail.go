package mail

import (
	"fmt"
	"html"
	"net/smtp"

	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gitlab.com/octomancer/packt-scraper/locations"
)

var initDone bool = false

func getConfig() {
	viper.SetConfigName(".mail")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		glog.Fatalf("Fatal error config file: %s \n", err)
	}
	initDone = true
}

func MakeEmail(subject, bookTitle, bookFeatures, imageURL string) string {
	return fmt.Sprintf("Subject: %s\nContent-type: text/html\n\n<html><body><a href=\"%s\">%s</a><br><br>%s<br>%s<br><img src=\"%s\"></body></html>",
		subject,
		locations.FreeLearningURL,
		locations.FreeLearningURL,
		bookTitle,
		bookFeatures,
		html.EscapeString(imageURL),
	)
}

// Send an email through my gmail account
func SendEmail(to []string, body string) error {
	if !initDone {
		getConfig()
	}
	err := smtp.SendMail(
		viper.GetString("mailhost"),
		smtp.PlainAuth(
			"",
			viper.GetString("mailuser"),
			viper.GetString("mailpass"),
			viper.GetString("maildomain"),
		),
		viper.GetString("mailuser"),
		to,
		[]byte(body),
	)
	return err
}
