package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type postgresPostVoteRepository struct {
	Conn *pgxpool.Pool
}

// NewPostgresPostVoteRepository will create an object that represent the post vote. Repository interface
func NewPostgresPostVoteRepository(Conn *pgxpool.Pool) models.PostVoteRepository {
	return &postgresPostVoteRepository{
		Conn: Conn,
	}
}

// Like post vote ...
func (pv *postgresPostVoteRepository) Like(ctx context.Context, postVote *models.PostVote) (err error) {
	return nil
}

// Dislike post vote ...
func (pv *postgresPostVoteRepository) Dislike(ctx context.Context, postVote *models.PostVote) (err error) {
	return nil
}

// Delete post vote ...
func (pv *postgresPostVoteRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}
