package entity

import "time"

type Contact struct {
	Id          string
	Name        string
	PhoneNumber string
	CreatedAt   time.Time
}
