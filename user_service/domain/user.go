package domain

import (
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrEmptyPassword = errors.New("empty password")
)

// User is user model
type User struct {
	UUID     string `json:"uuid"`
	Username string `json:"username" validate:"required,min=6,max=30"`
	Password string `json:"password" validate:"required,min=8,max=15"`
	Email    string `json:"email" validate:"required,min=3,max=250,email"`
	Bio      string `json:"bio"`
	Image    string `json:"image"`
}

// Validate validates fields of user model
func (u User) Validate(ctx context.Context) error {
	if err := utils.ValidateStruct(ctx, u); err != nil {
		return err
	}
	return nil
}

// HashPassword makes password field crypted
func (u *User) SetUUID() error {
	if u.UUID == "" {
		u.UUID = uuid.New().String()
	}
	return nil
}

// HashPassword makes password field crypted
func (u *User) HashPassword() error {
	if len(u.Password) == 0 {
		return ErrEmptyPassword
	}
	h, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.Wrap(err, "can't generate password")
	}
	u.Password = string(h)
	return nil
}

// CheckPassword checks user password correct
func (u *User) CheckPassword(plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plain))
	return err == nil
}

// HidePassword hide user password
func (u *User) HidePassword() {
	u.Password = ""
}
