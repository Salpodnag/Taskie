package repositories

import (
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

// func (upr *UserProjectRepository) AddUserToProject(UserID int, ProjectID int, role models.UserProjectRole){
// 	query :=
// }
