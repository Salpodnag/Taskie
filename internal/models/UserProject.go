package models

type UserProject struct {
	Id      int             `json:"int"`
	User    User            `json:"user"`
	Project Project         `json:"project"`
	Role    UserProjectRole `json:"role"`
}
