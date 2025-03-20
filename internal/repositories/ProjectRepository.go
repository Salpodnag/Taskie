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
	var owner models.User
	query := `
			SELECT p.id, p.name, p.created_at, u.id, u.email, u.username, u.time_registration
			FROM project p
			LEFT JOIN user_account u 
			ON p.owner_id = u.id
			WHERE p.id = $1`
	row := pr.db.QueryRow(context.Background(), query, id)

	err := row.Scan(&project.Id, &project.Name, &project.CreatedAt, &owner.Id, &owner.Email, &owner.Username, &owner.TimeRegistration)
	if err != nil {
		return nil, fmt.Errorf("failed to get project by id : %w", err)
	}
	project.Owner = owner
	return &project, nil
}

func (pr *ProjectRepository) GetAllProjects(id int) ([]models.Project, error) {
	projects := make([]models.Project, 0)
	query := `
			SELECT p.id, p.name, p.created_at, u.id, u.email, u.username, u.time_registration
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

func (pr *ProjectRepository) DeleteProject(id int) error {
	query := `
		DELETE FROM project
		WHERE id = $1`
	_, err := pr.db.Exec(context.Background(), query, id)
	if err != nil {
		return fmt.Errorf("failed to delete project: %w", err)
	}
	return nil
}

func (pr *ProjectRepository) CreateDefaultRoles(projectId int) error {
	var roleId int
	query := `
	SELECT id
	FROM user_project_role
	WHERE project_id = $1 AND name='Участник'`
	err := pr.db.QueryRow(context.Background(), query, projectId).Scan(&roleId)
	if err == pgx.ErrNoRows {
		query := `INSERT INTO user_project_role(project_id, name)
				VALUES ($1, 'Участник')`
		_, err := pr.db.Exec(context.Background(), query, projectId)
		if err != nil {
			return fmt.Errorf("Failed to insert role 'Участник': %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("Failed to check if role 'Участник' exists: %w", err)
	}

	query = `
	SELECT id
	FROM user_project_role
	WHERE project_id = $1 AND name='Владелец'`
	err = pr.db.QueryRow(context.Background(), query, projectId).Scan(&roleId)
	if err == pgx.ErrNoRows {
		query := `INSERT INTO user_project_role(project_id, name)
				VALUES ($1, 'Владелец')`
		_, err := pr.db.Exec(context.Background(), query, projectId)
		if err != nil {
			return fmt.Errorf("Failed to insert role 'Владелец': %w", err)
		}
	} else if err != nil {
		return fmt.Errorf("Failed to check if role 'Владелец' exists: %w", err)
	}
	return nil
}

// func (pr *ProjectRepository) AddUserToProject(projectId int, userId int)
