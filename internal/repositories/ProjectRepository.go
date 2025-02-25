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

func (pr *ProjectRepository) CreateProject(project models.Project) error {
	_, err := pr.db.Exec(context.Background(), "INSERT INTO project(name, owner_id, createdAt) VALUES ($1, $2, $3)",
		project.Name, project.Owner.Id, project.CreatedAt)
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
