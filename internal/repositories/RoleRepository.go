package repositories

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RoleRepository struct {
	db *pgxpool.Pool
}

func NewRoleRepository(db *pgxpool.Pool) *RoleRepository {
	return &RoleRepository{
		db: db,
	}
}

func (rl *RoleRepository) CreateDefaultRole(ProjectID uuid.UUID, role string) error {
	var roleId int
	query := `
	SELECT id
	FROM user_project_role
	WHERE project_id = $1 AND name=$2`
	err := rl.db.QueryRow(context.Background(), query, ProjectID, role).Scan(&roleId)
	if err == pgx.ErrNoRows {
		query := `INSERT INTO user_project_role(project_id, name)
				VALUES ($1, $2)`
		_, err := rl.db.Exec(context.Background(), query, ProjectID, role)
		if err != nil {
			return fmt.Errorf("failed to insert role: %w", err, role)
		}
	} else if err != nil {
		return fmt.Errorf("failed to check if role 'Участник' exists: %w", err, "role", role)
	}

	return nil
}
