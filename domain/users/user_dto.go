package users

import (
	"strings"

	"github.com/thanhftu/bookstore_users-api/utils/errors"
)

const (
	// StatusActive have active status of user
	StatusActive = "active"
)

// User struct
type User struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Users is slice of single user
type Users []User

// Validate user input before saving
func (user *User) Validate() *errors.RestErr {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	if user.Email == "" {

		return errors.NewBadRequestError("invalid email adress")
	}
	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {

		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
