package adapters

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/library_service/domain/article"
)

type GORMConfig struct {
	Type string
	User string
	Pass string
	Name string
	Host string
	Port string
}

type GORMArticleModel struct {
	gorm.Model
	UUID        string `json:"uuid"`
	UserUUID    string `json:"user_uuid"`
	Title       string `json:"title"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func (m *GORMArticleModel) protoDomainArticle() *article.Article {
	return &article.Article{
		UUID:        m.UUID,
		UserUUID:    m.UserUUID,
		Title:       m.Title,
		Slug:        m.Slug,
		Description: m.Description,
		Body:        m.Body,
	}
}

func (m *GORMArticleModel) protoGORMArticle(article *article.Article) {
	m.UUID = article.UUID
	m.UserUUID = article.UserUUID
	m.Title = article.Title
	m.Slug = article.Slug
	m.Description = article.Description
	m.Body = article.Body
}

type GORMArticleRepository struct {
	DB *gorm.DB
}

func NewGORMArticleRepository(config GORMConfig) (*GORMArticleRepository, error) {
	DBString := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	URL := fmt.Sprintf(DBString, config.User, config.Pass, config.Host, config.Port, config.Name)
	db, err := gorm.Open(config.Type, URL)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect database")
	}
	if err := migrate(db); err != nil {
		return nil, errors.Wrap(err, "cannot migrate database")
	}
	return &GORMArticleRepository{DB: db}, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&GORMArticleModel{},
	).Error
}

func (r *GORMArticleRepository) Create(ctx context.Context, article *article.Article) error {
	gormArticle := &GORMArticleModel{}
	gormArticle.protoGORMArticle(article)
	err := r.DB.Create(gormArticle).Error
	if err != nil {
		return errors.Wrap(err, "cannot create article")
	}
	return nil
}
func (r *GORMArticleRepository) GetBySlug(ctx context.Context, slug string) (*article.Article, error) {
	var gormArticle GORMArticleModel
	if err := r.DB.Where(GORMArticleModel{Slug: slug}).First(&gormArticle).Error; err != nil {
		return nil, errors.Wrap(err, "cannot find user")
	}
	return gormArticle.protoDomainArticle(), nil
}
