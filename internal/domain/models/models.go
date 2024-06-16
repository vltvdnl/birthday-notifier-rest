package models

import "time"

type User struct {
	ID            int64
	First_name    string
	Last_name     string
	Email         string
	// Password      strings
	Birthdate     time.Time
	Subscriptions []User
}
