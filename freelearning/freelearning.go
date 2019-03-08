package freelearning

import (
	"encoding/json"
	"fmt"
	"html"
	"os"
	"time"

	"github.com/golang/glog"
	"github.com/nlopes/slack"
	"github.com/spf13/viper"
	"gitlab.com/octomancer/packt-scraper/books"
	"gitlab.com/octomancer/packt-scraper/client"
	"gitlab.com/octomancer/packt-scraper/locations"
	"gitlab.com/octomancer/packt-scraper/mail"
	slac "gitlab.com/octomancer/packt-scraper/slack"
)

// Send an email about today's free learning ebook
func (f *FreeLearningBooks) SendEmail() {
	// Get HTTP REST client
	api := client.GetClient()

	// Get midnight of today and tomorrow
	t := time.Now()
	t1 := t.Format("2006-01-02")
	t = t.Add(time.Hour * 24)
	t2 := t.Format("2006-01-02")

	// URL to get today's free learning JSON
	url := fmt.Sprintf("%sfree-learning-v1/offers?dateFrom=%s&dateTo=%s", locations.ServicesURL, t1, t2)
	api.GetPageJSON(url, f)

	if f.Count == 0 {
		glog.V(0).Infof("No free book at %s\n", url)
		return
	}

	fb := f.Data[0]
	fb.Summary = &FreeLearningBookSummary{}
	glog.V(1).Infof("Got free book")
	fb.GetSummary()
	glog.V(1).Infof("Got summary")
	fb.PostToSlack()
	glog.V(1).Infof("Sent slack post")

	// Get the list of people to send the email to
	email_recipients := viper.GetStringSlice("recipients")

	// Check if we've got this book
	eb := &books.EntitlementBooks{}
	eb.ParseMyEbooks()
	b := eb.FindBookWithProductId(fb.ProductId)

	email_body := ""
	if b == nil {
		glog.V(0).Infof("Claiming book id %s", fb.Summary.Title)
		email_body = mail.MakeEmail("New packt ebook to claim", fb.Summary.Title, fb.Summary.Features, fb.Summary.CoverImage)
	} else {
		glog.V(0).Infof("Not claiming book id %s - %s", b.ProductId, b.ProductName)
		email_body = mail.MakeEmail("Already have today's packt ebook", fb.Summary.Title, fb.Summary.Features, fb.Summary.CoverImage)
	}
	err := mail.SendEmail(email_recipients, email_body)
	if err == nil {
		glog.V(0).Infof("Email sent")
	} else {
		glog.V(0).Infof("Error sending email: %s", err)
	}
}

// Get the FLBS for a single book and save it to disk
func (f *FreeLearningBook) GetSummary() {
	// Get HTTP REST client
	api := client.GetClient()

	// Get the book info
	url := fmt.Sprintf("%sproducts/%s/summary", locations.StaticURL, f.ProductId)
	api.GetPageJSON(url, f.Summary)
	glog.V(1).Infof("Got summary JSON")

	summaryFile := locations.GetFlbsFilename(f.ProductId)
	glog.V(1).Infof("Got FLBS filename %s\n", summaryFile)

	fh, err := os.Create(summaryFile)
	if err != nil {
		glog.Fatalf("Error opening %s for writing: %s", summaryFile, err)
	}
	defer fh.Close()

	byts, err := json.Marshal(f)
	if err != nil {
		glog.Fatalf("Error marshalling books: %s\n", err)
	}

	_, err = fh.Write(byts)
	if err != nil {
		glog.Fatalf("Error writing to %s: %s", summaryFile, err)
	}
}

func (f *FreeLearningBook) PostToSlack() {
	text := slack.MsgOptionText(fmt.Sprintf("Today's free eBook from Packt Publishing is *%s*! <https://www.packtpub.com/packt/offers/free-learning|Click here to get it>.", f.Summary.Title), false)
	attachment := slack.Attachment{
		Text:      f.Summary.OneLiner,
		Title:     f.Summary.Title,
		TitleLink: "https://www.packtpub.com/packt/offers/free-learning",
		ThumbURL:  html.EscapeString(f.Summary.CoverImage),
	}
	slac.PostToSlack(text, attachment)
}
