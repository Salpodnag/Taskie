package repositories

import (
	"Taskie/internal/models"
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProjectRepository struct {
	db *pgxpool.Pool
}

func NewProjectRepository(db *pgxpool.Pool) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (pr *ProjectRepository) CreateProject(project *models.Project) error {
	query := `
        INSERT INTO project (id, name, description, color, privacy, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
        RETURNING id`
	err := pr.db.QueryRow(context.Background(), query, project.Id, project.Name, project.Description, project.Color, project.Privacy, project.CreatedAt).Scan(&project.Id)
	if err != nil {
		return fmt.Errorf("failed to insert project: %w", err)
	}
	return nil
}

func (pr *ProjectRepository) GetProjectById(projectID uuid.UUID) (*models.Project, error) {
	var project models.Project
	query := `
			SELECT p.id, p.name, p.description, p.color, p.privacy, p.created_at
			FROM project p
			WHERE p.id = $1`
	row := pr.db.QueryRow(context.Background(), query, projectID)

	err := row.Scan(&project.Id, &project.Name, &project.Description, &project.Color, &project.Privacy, &project.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to get project by id : %w", err)
	}
	return &project, nil
}

func (pr *ProjectRepository) GetAllProjects() ([]models.Project, error) {
	projects := make([]models.Project, 0)
	query := `
			SELECT p.id, p.name, p.description, p.color, p.privacy, p.created_at
			FROM project p
			`
	rows, err := pr.db.Query(context.Background(), query)
	if err != nil {
		return nil, fmt.Errorf("failed to get all projects: %w", err)
	}
	for rows.Next() {
		var project models.Project
		err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.Color, &project.Privacy, &project.CreatedAt)
		if err != nil {
			slog.Error("???: %w", err)
			return nil, fmt.Errorf("failed to get all projects: %w", err)
		}
		projects = append(projects, project)
	}

	return projects, nil
}

func (pr *ProjectRepository) DeleteProject(projectID uuid.UUID) error {
	query := `
		DELETE FROM project
		WHERE id = $1`
	_, err := pr.db.Exec(context.Background(), query, projectID)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}

// func (pr *ProjectRepository) AddUserToProject(projectId int, userId int)
