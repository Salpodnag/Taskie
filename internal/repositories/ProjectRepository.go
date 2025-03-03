package repositories

import (
	"Taskie/internal/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type ProjectRepository struct {
	db *pgx.Conn
}

func NewProjectRepository(db *pgx.Conn) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (pr *ProjectRepository) CreateProject(project *models.Project) error {
	query := `
        INSERT INTO project (name, created_at, owner_id)
        VALUES ($1, $2, $3)
        RETURNING id`
	err := pr.db.QueryRow(context.Background(), query, project.Name, project.CreatedAt, project.Owner.Id).Scan(&project.Id)
	if err != nil {
		return fmt.Errorf("failed to insert project: %w", err)
	}
	return nil
}

func (pr *ProjectRepository) GetProjectById(id int) (*models.Project, error) {
	var user models.User
	var project models.Project
	query := "SELECT p.id, p.name, p.created_at, ua.id, ua.email, ua.username, ua.password, ua.time_registration from project  p  join user_account  ua ON p.owner_id = ua.id WHERE p.id = $1"
	row := pr.db.QueryRow(context.Background(), query, id)
	err := row.Scan(&project.Id, &project.Name, &project.CreatedAt, &user.Id, &user.Email, &user.Username, &user.Password, &user.TimeRegistration)
	if err != nil {
		return nil, fmt.Errorf("failed to ged project by id: %w", err)
	}
	project.Owner = user
	return &project, nil
}
