package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"github.com/golang/glog"
	"github.com/spf13/viper"
	"gitlab.com/octomancer/packt-scraper/locations"
	"golang.org/x/net/publicsuffix"
)

var client *httpClient

func GetClient() *httpClient {
	if client == nil {
		client = newClient()
	}
	return client
}

// Get an http client with a cookie jar
func newClient() *httpClient {
	// Create cookie jar
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		glog.Fatalf("Error creating cookie jar: %s", err)
	}

	// Create HTTP client
	client = &httpClient{loggedIn: false, accessToken: ""}
	client.httpClient = &http.Client{Jar: jar}

	return client
}

func (c *httpClient) SetRedirect() {
	c.httpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
}

func (c *httpClient) ClearRedirect() {
	c.httpClient.CheckRedirect = nil
}

func (c *httpClient) Login() {
	if c.loggedIn {
		return
	}

	creds := &LoginRequest{
		Username: viper.GetString("email"),
		Password: viper.GetString("password"),
	}
	credBytes, err := json.Marshal(creds)
	if err != nil {
		glog.Fatalf("Error creating HTTP request: %s", err)
	}

	// Log in
	res, err := c.httpClient.Post(
		fmt.Sprintf("%s%s", locations.ServicesURL, locations.LoginPath),
		"application/json",
		bytes.NewReader(credBytes),
	)
	if err != nil {
		glog.Fatalf("Error logging in: %s", err)
	}
	defer res.Body.Close()
	lrs := &LoginResponses{}
	err = json.NewDecoder(res.Body).Decode(lrs)
	if err != nil {
		glog.Fatalf("Error decoding login JSON: %s", err)
	}

	c.accessToken = lrs.Data.Access
	c.loggedIn = true
	glog.V(0).Info("Logged in")
}

// Get a page from packtpub.com
func (c *httpClient) GetPageJSON(url string, obj interface{}) {
	// Get page and unmarshall it into obj
	glog.V(0).Infof("fetching %s", url)
	res, err := c.Do(url)
	if err != nil {
		glog.Fatalf("Error fetching url %s: %s\n", url, err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(obj)
	if err != nil {
		glog.Fatalf("Error unmarshalling free learning JSON: %s\n", err)
	}
}

func (c *httpClient) Do(url string) (*http.Response, error) {
	// Request page
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		glog.Fatalf("Error creating request: %s", err)
	}
	if c.loggedIn {
		req.Header.Add("Authorization", "Bearer "+c.accessToken)
	}
	return c.httpClient.Do(req)
}
