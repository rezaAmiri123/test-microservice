package pg

import (
	"context"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
	"github.com/stretchr/testify/require"
)

func TestPGArticleRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlDB")
	t.Run("Create", func(t *testing.T) {
		a := &article.Article{
			Title:       "article_create_title",
			Slug:        "article_create_slug",
			Body:        "article_create_body",
			Description: "article_create_description",
		}
		err := a.SetUUID(uuid.New().String())
		require.NoError(t, err)
		rows := sqlmock.NewRows([]string{"uuid", "user_uuid", "title", "slug", "description", "body", "created_at", "updated_at"}).AddRow(
			a.UUID,
			a.UserUUID,
			a.Title,
			a.Slug,
			a.Description,
			a.Body,
			time.Now(),
			time.Now(),
		)
		mock.ExpectQuery(createArticle).WithArgs(
			a.UUID,
			a.UserUUID,
			a.Title,
			a.Slug,
			a.Description,
			a.Body,
			//time.Now(),
			//time.Now(),
		).WillReturnRows(rows)
		repo := NewPGArticleRepository(sqlxDB)
		err = repo.Create(context.Background(), a)
		require.NoError(t, err)
	})
}

func TestPGArticleRepository_GetBySlug(t *testing.T) {
	t.Parallel()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlDB")
	t.Run("GetBySlug", func(t *testing.T) {
		a := &article.Article{
			UUID:        uuid.New().String(),
			Title:       "article_create_title",
			Slug:        "article_create_slug",
			Body:        "article_create_body",
			Description: "article_create_description",
		}
		err := a.SetUUID(uuid.New().String())
		require.NoError(t, err)
		rows := sqlmock.NewRows([]string{"uuid", "user_uuid", "title", "slug", "description", "body", "created_at", "updated_at"}).AddRow(
			a.UUID,
			a.UserUUID,
			a.Title,
			a.Slug,
			a.Description,
			a.Body,
			time.Now(),
			time.Now(),
		)
		mock.ExpectQuery(getArticleBySlug).WithArgs(a.Slug).WillReturnRows(rows)
		repo := NewPGArticleRepository(sqlxDB)
		resultArticle, err := repo.GetBySlug(context.Background(), a.Slug)
		require.NoError(t, err)
		require.Equal(t, a.UserUUID, resultArticle.UserUUID)
	})
}
