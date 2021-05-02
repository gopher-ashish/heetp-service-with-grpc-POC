package model

import (
	"regexp"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// BeforeCreate validates the model before creating in the database.
func (u User) BeforeCreate(scope *gorm.Scope) (err error) {
	namespace, _ := uuid.FromString("user")
	userID := uuid.NewV5(namespace, uuid.NewV4().String()+u.FirstName+u.Username+time.Now().String())
	scope.SetColumn("user_id", userID.String())

	if err := u.Validate(); err != nil {
		return err
	}

	return nil
}

// Validate owner - this should simply validates that the fields are valid, not any username/email
func (u *User) Validate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.Username, validation.Required, validation.Length(3, 255)),
		validation.Field(&u.Email, validation.Required, validation.Length(5, 255), validation.Match(
			regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)). // Accept 20-byte or 32-byte addresses
			Error("invalid email"),
		),
		validation.Field(&u.FirstName, validation.Required, validation.Length(3, 255), validation.Match(
			regexp.MustCompile(`^[a-zA-Z0-9 ,.'-]+$`)). // Accept 20-byte or 32-byte addresses
			Error("invalid name"),
		))
}
