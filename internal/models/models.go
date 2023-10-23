package models

import "time"

// User - holds information for a user
type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Position string `json:"position"`
}

// Timesheet - holds information for a users timesheet
type Timesheet struct {
	Id            int            `json:"id"`
	UserId        int            `json:"user_id"`
	DateWeek      time.Time      `json:"date_week"`
	DateSubmitted time.Time      `json:"date_submitted"`
	Data          map[string]any `json:"data"`
}

// ErrorResponse - holds error information for response
type ErrorResponse struct {
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
}
