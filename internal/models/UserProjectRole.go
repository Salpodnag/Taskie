package models

type UserProjectRole struct {
	Id         int    `json:"id"`
	Project_id int    `json:"projectId"`
	Name       string `json:"name"`
}
