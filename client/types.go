package client

import "net/http"

type httpClient struct {
	loggedIn    bool
	accessToken string
	httpClient  *http.Client
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Refresh string `json:"refresh"`
	Access  string `json:"access"`
}

type LoginResponses struct {
	Data LoginResponse `json:"data"`
}
