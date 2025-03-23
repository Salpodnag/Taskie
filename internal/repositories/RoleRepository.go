package repositories

import (
	"context"
	"fmt"

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

func (rl *RoleRepository) CreateDefaultRoles(rlojectId int) error {
	var roleId int
	query := `
	SELECT id
	FROM user_project_role
	WHERE rloject_id = $1 AND name='Участник'`
	err := rl.db.QueryRow(context.Background(), query, rlojectId).Scan(&roleId)
	if err == pgx.ErrNoRows {
		query := `INSERT INTO user_project_role(project_id, name)
				VALUES ($1, 'Участник')`
		_, err := rl.db.Exec(context.Background(), query, rlojectId)
		if err != nil {
			return fmt.Errorf("Failed to insert role 'Участник': %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("Failed to check if role 'Участник' exists: %w", err)
	}

	query = `
	SELECT id
	FROM user_project_role
	WHERE rloject_id = $1 AND name='Владелец'`
	err = rl.db.QueryRow(context.Background(), query, rlojectId).Scan(&roleId)
	if err == pgx.ErrNoRows {
		query := `INSERT INTO user_project_role(project_id, name)
				VALUES ($1, 'Владелец')`
		_, err := rl.db.Exec(context.Background(), query, rlojectId)
		if err != nil {
			return fmt.Errorf("Failed to insert role 'Владелец': %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("Failed to check if role 'Владелец' exists: %w", err)
	}
	return nil
}
