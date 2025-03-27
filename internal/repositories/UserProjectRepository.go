package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (upr *UserProjectRepository) AddUserToProject(userID uuid.UUID, projectID uuid.UUID, roleID int) error {
	query := `INSERT INTO user_project (user_id, project_id, user_project_role_id)
	values ($1, $2, $3)
	`
	_, err := upr.db.Exec(context.Background(), query, userID, projectID, roleID)
	if err != nil {
		return fmt.Errorf("failed to add user to project (userID=%s, projectID=%s, roleID=%d): %w",
			userID, projectID, roleID, err)
	}
	return nil

}
