package adapters

import (
	"context"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/user_service/domain"
)

// User is user model
type GORMUserModel struct {
	gorm.Model
	UUID     string `json:"uuid"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

type GORMConfig struct {
	Type string
	User string
	Pass string
	Name string
	Host string
	Port string
}

func (m *GORMUserModel) ProtoDomainUser() *domain.User {
	return &domain.User{
		UUID:     m.UUID,
		Username: m.Username,
		Password: m.Password,
		Email:    m.Email,
		Bio:      m.Bio,
		Image:    m.Image,
	}
}

func (m *GORMUserModel) protoGORMUser(user *domain.User) {
	m.UUID = user.UUID
	m.Username = user.Username
	m.Password = user.Password
	m.Email = user.Email
	m.Bio = user.Bio
	m.Image = user.Image
}

type GORMUserRepository struct {
	DB *gorm.DB
}

func NewGORMUserRepository(config GORMConfig) (*GORMUserRepository, error) {
	DBString := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	URL := fmt.Sprintf(DBString, config.User, config.Pass, config.Host, config.Port, config.Name)
	db, err := gorm.Open(config.Type, URL)
	if err != nil {
		return nil, errors.Wrap(err, "cannot connect database")
	}
	if err := migrate(db); err != nil {
		return nil, errors.Wrap(err, "cannot migrate database")
	}
	return &GORMUserRepository{DB: db}, nil
}

func (r *GORMUserRepository) Create(ctx context.Context, user *domain.User) error {
	gormUser := &GORMUserModel{}
	gormUser.protoGORMUser(user)
	r.DB.LogMode(true)
	err := r.DB.Create(gormUser).Error
	if err != nil {
		return errors.Wrap(err, "cannot create user")
	}
	return nil
}
func (r *GORMUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	var gormUser GORMUserModel
	if err := r.DB.Where(GORMUserModel{Username: username}).First(&gormUser).Error; err != nil {
		return nil, errors.Wrap(err, "cannot find user")
	}
	return gormUser.ProtoDomainUser(), nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&GORMUserModel{},
	).Error
}
