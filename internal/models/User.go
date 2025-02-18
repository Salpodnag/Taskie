package models

import "time"

type User struct {
	Id               int       `json:"id"`
	Email            string    `json:"email"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	TimeRegistration time.Time `json:"timeRegistration"`
}
