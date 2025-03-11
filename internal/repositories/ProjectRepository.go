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
	var project models.Project
	query := `
			SELECT id, name, created_at, owner_id 
			FROM project 
			WHERE id = $1`
	row := pr.db.QueryRow(context.Background(), query, id)
	err := row.Scan(&project.Id, &project.Name, &project.CreatedAt, &project.Owner.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to get project by id: %w", err)
	}
	return &project, nil
}

func (pr *ProjectRepository) GetAllProjects(id int) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	query := `
			SELECT p.id, p.name, p.created_at, u.name, u.email, u.username, u.time_registration
			FROM project p
			LEFT JOIN user_account u 
			ON p.owner_id = u.id
			WHERE u.id = $1`
	rows, err := pr.db.Query(context.Background(), query, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get all projects: %w", err)
	}
	for rows.Next() {
		var project models.Project
		var owner models.User
		err := rows.Scan(&project.Id, &project.Name, &project.CreatedAt, &owner.Id, &owner.Email, &owner.Username, &owner.TimeRegistration)
		if err != nil {
			return nil, fmt.Errorf("failed to get all projects: %w", err)
		}
		project.Owner = owner
		projects = append(projects, project)
	}

	return projects, nil
}
