package model

import (
	"github.com/internal/connection"
)

type User struct {
	Id        uint   `gorm:"primary_key" json:"-"`
	UserID    string `gorm:"not null;unique" json:"userID"`
	Username  string `gorm:"not null" json:"username"`
	FirstName string `gorm:"not null" json:"name"`
	Email     string `gorm:"not null" json:"email"`
}

type UserRepository interface {
	Create(User) (User, error)
	ShowUsers() (users []User, err error)
}

// Create a new owner.
func (u *User) Create(user User) (User, error) {
	if err := connection.Get().Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil

}

//get all existing users
func (u *User) ShowUsers() (users []User, err error) {
	if err := connection.Get().Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func NewUserReository() UserRepository {
	return &User{}
}
