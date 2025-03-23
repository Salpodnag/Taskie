package repositories

import (
	"Taskie/internal/models"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) CreateUser(user models.User) error {
	err := ur.db.QueryRow(context.Background(), "INSERT INTO user_account(id, email, username, password, time_registration) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		user.Id, user.Email, user.Username, user.Password, user.TimeRegistration).Scan(&user.Id)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}
	return nil
}

func (ur *UserRepository) GetUserByEmailOrUsername(name string) (*models.User, error) {
	var user models.User

	query := `SELECT id, email, username, password, time_registration FROM user_account where email=$1 OR username=$1`
	row := ur.db.QueryRow(context.Background(), query, name)
	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.Password, &user.TimeRegistration)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found with email/username %w", err)
		}
		return nil, fmt.Errorf("failed to get user by email/username %w", err)
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := `
			SELECT id, email, username, time_registration 
			FROM user_account 
			WHERE email=$1 OR username=$1`
	row := ur.db.QueryRow(context.Background(), query, email)
	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.TimeRegistration)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found with email %w", err)
		}
		return nil, fmt.Errorf("failed to get user by email %w", err)
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	query := `
			SELECT id, email, username, time_registration 
			FROM user_account 
			WHERE email=$1 OR username=$1`
	row := ur.db.QueryRow(context.Background(), query, username)
	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.TimeRegistration)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found with username %w", err)
		}
		return nil, fmt.Errorf("failed to get user by username %w", err)
	}
	return &user, nil
}

func (ur *UserRepository) GetUserById(UserID uuid.UUID) (*models.User, error) {

	var user models.User

	query := `
			SELECT id, email, username, time_registration 
			FROM user_account 
			WHERE  id=$1`
	row := ur.db.QueryRow(context.Background(), query, UserID)
	err := row.Scan(&user.Id, &user.Email, &user.Username, &user.TimeRegistration)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("user not found with id %w", err)
		}
		return nil, fmt.Errorf("failed to get user by id %w", err)
	}
	return &user, nil
}
