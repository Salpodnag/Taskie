package repositories

import (
	"Taskie/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserProjectRepository struct {
	db *pgxpool.Pool
}

func NewUserProjectRepository(db *pgxpool.Pool) *UserProjectRepository {
	return &UserProjectRepository{
		db: db,
	}
}

func (upr *UserProjectRepository) AddUserToProject(userID int, projectID int, role models.UserProjectRole) error {
	query := `INSERT INTO user_project_role (user_id, project_id, user_project_role_id)
	values ($1, $2, $3)
	RETURNING 1
	`
	_, err := upr.db.Exec(context.Background(), query, userID, projectID, role)
	if err != nil {
		return fmt.Errorf("failed to add user to project: %w", err, "userId:", userID, "projectID", projectID, "role", role)
	}
	return nil

}
