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

func (rl *RoleRepository) GetRoleID(projectID uuid.UUID, role string) (int, error) {
	var roleID int
	query := `SELECT id FROM user_project_role WHERE project_id = $1 AND name = $2`
	err := rl.db.QueryRow(context.Background(), query, projectID, role).Scan(&roleID)
	if err == pgx.ErrNoRows {
		return 0, nil
	} else if err != nil {
		return 0, fmt.Errorf("failed to check role: %w", err)
	}
	return roleID, nil
}

func (rl *RoleRepository) CreateRole(projectID uuid.UUID, role string) (int, error) {
	var roleID int
	query := `INSERT INTO user_project_role (project_id, name) VALUES ($1, $2) RETURNING id`
	err := rl.db.QueryRow(context.Background(), query, projectID, role).Scan(&roleID)
	if err != nil {
		return 0, fmt.Errorf("failed to insert role: %w", err)
	}
	return roleID, nil
}

func (rl *RoleRepository) GetOrCreateDefaultRole(projectID uuid.UUID, role string) (int, error) {
	roleID, err := rl.GetRoleID(projectID, role)
	if err != nil {
		return 0, err
	}
	if roleID != 0 {
		return roleID, nil
	}
	return rl.CreateRole(projectID, role)
}
