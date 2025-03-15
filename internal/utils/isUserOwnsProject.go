package utils

import "Taskie/internal/models"

func isOwner(project models.Project, user models.User) bool {
	return project.Owner == user
}
