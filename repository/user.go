package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type postgresUserRepository struct {
	Conn *pgxpool.Pool
}

// NewPostgresUserRepository will create an object that represent the user. Repository interface
func NewPostgresUserRepository(Conn *pgxpool.Pool) models.UserRepository {
	return &postgresUserRepository{
		Conn: Conn,
	}
}

// Create user ...
func (u *postgresUserRepository) Create(ctx context.Context, user *models.User) (id int, err error) {
	return 0, nil
}

// Update user ...
func (u *postgresUserRepository) Update(ctx context.Context, user *models.User) (err error) {
	return nil
}

// GetByID user ...
func (u *postgresUserRepository) GetByID(ctx context.Context, id int) (user *models.User, err error) {
	return nil, nil
}

// GetByEmail user ...
func (u *postgresUserRepository) GetByEmail(ctx context.Context, email string) (user *models.User, err error) {
	return nil, nil
}

// Delete user ...
func (u *postgresUserRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}
