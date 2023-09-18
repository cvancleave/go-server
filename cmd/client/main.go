package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"go-server/internal/config"
	"go-server/internal/models"
	"go-server/internal/utils"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type client struct {
	serverUrl    string
	serverConfig *config.Config
	http         *http.Client
}

func main() {

	// load server config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("error loading config")
	}

	// create handler
	client := &client{
		serverUrl:    "http://localhost:4000",
		serverConfig: cfg,
		http:         &http.Client{},
	}

	// create jwt token for authentication
	token, err := utils.NewToken("userIDHere", cfg.TokenKey, cfg.TokenIssuer, cfg.TokenTimeout)
	if err != nil {
		log.Fatal(err)
	}

	// use each endpoint
	client.getUser(token)
	client.postUser(token)
}

// getUser - sends an authenticated get request to the /user endpoint
func (c *client) getUser(token string) {

	endpoint := c.serverUrl + "/user?id=1"

	// setup request and add auth header
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// send request
	resp, err := c.http.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// unmarshal response into map to account for possible error message
	body := map[string]any{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Fatal(err)
	}

	// check response code
	if resp.StatusCode != 200 {
		err = errors.New(body["error_message"].(string))
		log.Fatal(err)
	}

	log.Info(body)
}

// postUser - sends an authenticated post request to the /user endpoint
func (c *client) postUser(token string) {

	endpoint := c.serverUrl + "/user"

	user := &models.User{Name: "userNameHere"}
	userJson, _ := json.Marshal(user)

	// setup request and add auth header
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(userJson))
	req.Header.Set("Authorization", "Bearer "+token)

	// send request
	resp, err := c.http.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// unmarshal response
	body := map[string]any{}
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		log.Fatal(err)
	}

	// check response code
	if resp.StatusCode != 200 {
		err = errors.New(body["error_message"].(string))
		log.Fatal(err)
	}

	log.Info(body)
}
