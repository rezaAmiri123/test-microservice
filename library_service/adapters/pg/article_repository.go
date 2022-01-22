package pg

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
)

type PGArticleModel struct {
	UUID        string    `json:"uuid" db:"uuid"`
	UserUUID    string    `json:"user_uuid" db:"user_uuid"`
	Title       string    `json:"title" db:"title"`
	Slug        string    `json:"slug" db:"slug"`
	Description string    `json:"description" db:"description"`
	Body        string    `json:"body" db:"body"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewPGArticleRepository(db *sqlx.DB) *PGArticleRepository {
	return &PGArticleRepository{DB: db}
}

// News Repository
type PGArticleRepository struct {
	DB *sqlx.DB
}

func (m *PGArticleModel) protoDomainArticle() *article.Article {
	return &article.Article{
		UUID:        m.UUID,
		UserUUID:    m.UserUUID,
		Title:       m.Title,
		Slug:        m.Slug,
		Description: m.Description,
		Body:        m.Body,
	}
}

func (m *PGArticleModel) protoPGArticle(article *article.Article) {
	m.UUID = article.UUID
	m.UserUUID = article.UserUUID
	m.Title = article.Title
	m.Slug = article.Slug
	m.Description = article.Description
	m.Body = article.Body
}

func (r *PGArticleRepository) Create(ctx context.Context, article *article.Article) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGArticleRepository.Create")
	defer span.Finish()

	pgArticle := &PGArticleModel{}
	pgArticle.protoPGArticle(article)

	if err := r.DB.QueryRowxContext(
		ctx,
		createArticle,
		pgArticle.UUID,
		pgArticle.UserUUID,
		pgArticle.Title,
		pgArticle.Slug,
		pgArticle.Description,
		pgArticle.Body,
	).Err(); err != nil {
		return errors.Wrap(err, "create article")
	}

	// gormArticle := &GORMArticleModel{}
	// gormArticle.protoGORMArticle(article)
	// err := r.DB.Create(gormArticle).Error
	// if err != nil {
	// 	return errors.Wrap(err, "cannot create article")
	// }
	return nil
}
func (r *PGArticleRepository) GetBySlug(ctx context.Context, slug string) (*article.Article, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGArticleRepository.GetBySlug")
	defer span.Finish()

	articleModel := &PGArticleModel{}
	if err := r.DB.GetContext(ctx, articleModel, getArticleBySlug, slug); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetNewsByID.GetContext")
	}
	return articleModel.protoDomainArticle(), nil
}
