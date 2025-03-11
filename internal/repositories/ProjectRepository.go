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

func (pr *ProjectRepository) GetAllProjects() ([]models.Project, error) {
	projects := []models.Project{}
	project := models.Project{}
	query := `
			SELECT id, name, created_at, owner_id
			FROM project`
	rows, err := pr.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all projects: %w", err)
	}
	for rows.Next() {
		rows.Scan(&project.Id, &project.Name, &project.CreatedAt, &project.Owner.Id)
		projects = append(projects, project)
	}

	return projects, nil
}
