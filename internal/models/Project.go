package models

import "time"

type Project struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Owner     User      `json:"user"`
	CreatedAt time.Time `json:"createdAt`
}
