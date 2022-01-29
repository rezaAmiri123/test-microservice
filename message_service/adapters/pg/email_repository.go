package pg

import (
	"context"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
	"time"
)

type PGEmailModel struct {
	UUID      string    `json:"uuid" db:"uuid"`
	From      string    `json:"from" db:"from"`
	To        []string  `json:"to" db:"to"`
	Subject   string    `json:"subject" db:"subject"`
	Body      string    `json:"body" db:"body"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

func NewPGEmailRepository(db *sqlx.DB) *PGEmailRepository {
	return &PGEmailRepository{DB: db}
}

// News Repository
type PGEmailRepository struct {
	DB *sqlx.DB
}

func (m *PGEmailModel) protoDomainEmail() *email.Email {
	return &email.Email{
		UUID:      m.UUID,
		From:      m.From,
		To:        m.To,
		Subject:   m.Subject,
		Body:      m.Body,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

func (m *PGEmailModel) protoPGEmail(e *email.Email) {
	m.UUID = e.UUID
	m.From = e.From
	m.To = e.To
	m.Subject = e.Subject
	m.Body = e.Body
}

func (r *PGEmailRepository) Create(ctx context.Context, e *email.Email) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGEmailRepository.Create")
	defer span.Finish()

	pgEmail := &PGEmailModel{}
	pgEmail.protoPGEmail(e)

	if err := r.DB.QueryRowxContext(
		ctx,
		createEmail,
		pgEmail.UUID,
		pgEmail.From,
		pgEmail.To,
		pgEmail.Subject,
		pgEmail.Body,
	).Err(); err != nil {
		return errors.Wrap(err, "create email")
	}

	return nil
}

func (r *PGEmailRepository) GetByUUID(ctx context.Context, uuid string) (*email.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "PGEmailRepository.GetBySlug")
	defer span.Finish()

	articleModel := &PGEmailModel{}
	if err := r.DB.GetContext(ctx, articleModel, getEmailByUUID, uuid); err != nil {
		return nil, errors.Wrap(err, "emailRepo.GetEmailByUUID.GetContext")
	}
	return articleModel.protoDomainEmail(), nil
}
