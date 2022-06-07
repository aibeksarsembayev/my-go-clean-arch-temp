package repository

import "github.com/jackc/pgx/v4/pgxpool"

type postgresUserRepository struct {
	Conn *pgxpool.Pool
}

// NewPostgresUserRepository will create an object that represent the user. Repository interface
func NewPostgresUserRepository(Conn *pgxpool.Pool) models.UserRepository {
	return &NewPostgresUserRepository{
		Conn: Conn,
	}
}
