package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-server/internal/config"
	"go-server/internal/models"
	"net/http"
	"net/url"
	"time"
)

type client struct {
	serverUrl    string
	serverConfig *config.Config
	http         *http.Client
}

// basic http client used for testing endpoints
func main() {

	// load server config
	cfg, err := config.Load()
	if err != nil {
		panic("error loading config")
	}

	// create handler
	client := &client{
		serverUrl:    "http://localhost:4000",
		serverConfig: cfg,
		http:         &http.Client{},
	}

	// test endpoints
	token := client.login()
	client.postTimesheet(token)
	client.getTimesheets(token)
}

// login - test POST /login
func (c *client) login() string {

	endpoint := c.serverUrl + "/login"

	formData := url.Values{
		"email":    {"test@email.com"},
		"password": {"testPassword"},
	}

	// send form
	resp, err := http.PostForm(endpoint, formData)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// get token response
	token := ""
	if err := json.NewDecoder(resp.Body).Decode(&token); err != nil {
		panic(err)
	}

	return token
}

func (c *client) postTimesheet(token string) {

	endpoint := c.serverUrl + "/timesheet"

	timesheet := &models.Timesheet{
		UserId:        1,
		DateWeek:      time.Now(),
		DateSubmitted: time.Now(),
	}
	timesheetJson, _ := json.Marshal(timesheet)

	// setup request and add auth header
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(timesheetJson))
	req.Header.Set("Authorization", "Bearer "+token)

	// send request
	resp, err := c.http.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// check response code
	if resp.StatusCode != 200 {
		panic(errors.New("post timesheet bad response code"))
	}

	// unmarshal response
	var success bool
	if err := json.NewDecoder(resp.Body).Decode(&success); err != nil {
		panic(err)
	}

	fmt.Println(success)
}

// getTimesheets - test GET /timesheet
func (c *client) getTimesheets(token string) {

	endpoint := c.serverUrl + "/timesheets?userId=1"

	// setup request and add auth header
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	// send request
	resp, err := c.http.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// check response code
	if resp.StatusCode != 200 {
		panic(errors.New("get timesheets bad response code"))
	}

	// unmarshal response into map to account for possible error message
	timesheets := []models.Timesheet{}
	if err := json.NewDecoder(resp.Body).Decode(&timesheets); err != nil {
		panic(err)
	}

	for _, t := range timesheets {
		fmt.Println(t.Id)
	}
}
