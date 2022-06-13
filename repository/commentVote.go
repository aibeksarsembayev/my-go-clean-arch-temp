package repository

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/quazar2000/my-go-clean-arch-temp/models"
)

type postgresCommentVoteRepository struct {
	Conn *pgxpool.Pool
}

// NewPostgresCommentVoteRepository will create an object tat represent the comment vote. Repository interface
func NewPostgresCommentVoteRepository(Conn *pgxpool.Pool) models.CommentVoteRepository {
	return &postgresCommentVoteRepository{
		Conn: Conn,
	}
}

// Like comment ...
func (cv *postgresCommentVoteRepository) Like(ctx context.Context, commentVote *models.CommentVote) (err error) {
	return nil
}

// Dislike comment ...
func (cv *postgresCommentVoteRepository) Dislike(ctx context.Context, commentVote *models.CommentVote) (err error) {
	return nil
}

// Delete comment vote ...
func (cv *postgresCommentVoteRepository) Delete(ctx context.Context, id int) (err error) {
	return nil
}
