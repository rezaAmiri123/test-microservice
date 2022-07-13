package pg

import (
	"context"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
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
	//).StructScan(pgArticle); err != nil {
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
		return nil, errors.Wrap(err, "PGArticleRepository.GetBySlug.GetContext")
	}
	return articleModel.protoDomainArticle(), nil
}

func (r *PGArticleRepository) List(ctx context.Context, query *pagnation.Pagination) (*article.ArticleList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGArticleRepository.List")
	defer span.Finish()

	var totalCount uint64
	if err := r.DB.QueryRowContext(ctx, totalArticleCountQuery).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "db.QueryRowContext")
	}
	if totalCount == 0 {
		return &article.ArticleList{Articles: []*article.Article{}}, nil
	}

	rows, err := r.DB.QueryxContext(ctx, articleListQuery, query.GetOffset(), query.GetLimit())
	if err != nil {
		return nil, errors.Wrap(err, "db.QueryContext")
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			err = errors.Wrap(closeErr, "rows.Close")
		}
	}()

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "rows.Err")
	}

	articles := make([]*article.Article, 0, query.GetSize())
	for rows.Next() {
		art := &article.Article{}
		if err := rows.Scan(
			&art.UUID,
			&art.UserUUID,
			&art.Title,
			&art.Slug,
			&art.Description,
			&art.Body,
		); err != nil {
			return nil, errors.Wrap(err, "rows.Scan")
		}
		articles = append(articles, art)
	}

	res := &article.ArticleList{}
	res.TotalCount = int64(totalCount)
	res.TotalPages = int64(query.GetTotalPages(int(totalCount)))
	res.Page = int64(query.GetPage())
	res.Size = int64(query.GetSize())
	res.HasMore = query.GetHasMore(int(totalCount))
	res.Articles = articles

	return res, err
}

// // FindEmailsByReceiver Find emails by receiver
//func (e *EmailsRepository) FindEmailsByReceiver(ctx context.Context, to string, query *utils.PaginationQuery) (list *models.EmailsList, err error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "EmailsRepository.FindEmailsByReceiver")
//	defer span.Finish()
//
//	var totalCount uint64
//	if err := e.db.QueryRowContext(ctx, totalCountQuery, to).Scan(&totalCount); err != nil {
//		return nil, errors.Wrap(err, "db.QueryRowContext")
//	}
//	if totalCount == 0 {
//		return &models.EmailsList{Emails: []*models.Email{}}, nil
//	}
//
//	rows, err := e.db.QueryxContext(ctx, findEmailByReceiverQuery, to, query.GetOffset(), query.GetLimit())
//	if err != nil {
//		return nil, errors.Wrap(err, "db.QueryxContext")
//	}
//	defer func() {
//		if closeErr := rows.Close(); closeErr != nil {
//			err = errors.Wrap(closeErr, "rows.Close")
//		}
//	}()
//
//	if err := rows.Err(); err != nil {
//		return nil, errors.Wrap(err, "rows.Err")
//	}
//
//	emails := make([]*models.Email, 0, query.GetSize())
//	for rows.Next() {
//		var mailTo string
//		email := &models.Email{}
//		if err := rows.Scan(
//			&email.EmailID,
//			&mailTo,
//			&email.From,
//			&email.Subject,
//			&email.Body,
//			&email.ContentType,
//			&email.CreatedAt,
//		); err != nil {
//			return nil, errors.Wrap(err, "rows.Scan")
//		}
//		email.SetToFromString(mailTo)
//		emails = append(emails, email)
//	}
//
//	return &models.EmailsList{
//		TotalCount: totalCount,
//		TotalPages: utils.GetTotalPages(totalCount, query.GetSize()),
//		Page:       query.GetPage(),
//		Size:       query.GetSize(),
//		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetSize()),
//		Emails:     emails,
//	}, err
//}
//
//
//
//
//	findEmailByReceiverQuery = `SELECT email_id, "to", "from", subject, body, content_type, created_at
//	FROM emails WHERE "to" ILIKE '%' || $1 || '%' ORDER BY created_at OFFSET $2 LIMIT $3`
//
//
//getProductByIdQuery = `SELECT p.product_id, p.name, p.description, p.price, p.created_at, p.updated_at
//	FROM products p WHERE p.product_id = $1`
//
//