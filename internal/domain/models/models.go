package models

import "time"

type User struct {
	First_name    string
	Last_name     string
	Email         string
	Birthdate     time.Time
	Subscriptions []User
}
