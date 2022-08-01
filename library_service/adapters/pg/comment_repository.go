package pg

import (
	"context"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
)

type PGCommentModel struct {
	UUID        string         `json:"uuid" db:"uuid"`
	UserUUID    string         `json:"user_uuid" db:"user_uuid"`
	ArticleUUID string         `json:"article_uuid" db:"article_uuid"`
	Article     PGArticleModel `json:"article"`
	Message     string         `json:"message" db:"message"`
	Likes       int64          `json:"likes" db:"likes"`
	CreatedAt   time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at" db:"updated_at"`
}

func (m *PGCommentModel) protoPGComment(comment *article.Comment) {
	m.UUID = comment.UUID
	m.UserUUID = comment.UserUUID
	m.ArticleUUID = comment.ArticleUUID
	m.Message = comment.Message
	m.Likes = comment.Likes
}

func (r *PGArticleRepository) CreateComment(ctx context.Context, comment *article.Comment) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGArticleRepository.CreateComment")
	defer span.Finish()

	pgComment := &PGCommentModel{}
	pgComment.protoPGComment(comment)

	if err := r.DB.QueryRowxContext(
		ctx,
		createComment,
		pgComment.UUID,
		pgComment.UserUUID,
		pgComment.ArticleUUID,
		pgComment.Message,
		pgComment.Likes,
	).Err(); err != nil {
		return errors.Wrap(err, "create comment")
	}
	return nil
}
