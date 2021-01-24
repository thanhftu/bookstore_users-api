package services

import (
	"fmt"

	"github.com/thanhftu/bookstore_users-api/domain/users"
	"github.com/thanhftu/bookstore_users-api/utils/cryptoutils"
	"github.com/thanhftu/bookstore_users-api/utils/dateutils"
	"github.com/thanhftu/bookstore_users-api/utils/errors"
)

// UsersService contains business logic relating to user
var (
	UsersService usersServiceInterface = &usersService{}
)

type usersService struct{}
type usersServiceInterface interface {
	GetUser(userID int64) (*users.User, *errors.RestErr)
	CreateUser(users.User) (*users.User, *errors.RestErr)
	UpdateUser(users.User) (*users.User, *errors.RestErr)
	DeleteUser(int64) *errors.RestErr
	SearchUser(string) (users.Users, *errors.RestErr)
	LoginUser(users.LoginRequest) (*users.User, *errors.RestErr)
}

// GetUser return a user
func (s *usersService) GetUser(userID int64) (*users.User, *errors.RestErr) {
	result := &users.User{ID: userID}
	if err := result.GET(); err != nil {
		return nil, err
	}
	return result, nil

}

// CreateUser create new user
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.Password = cryptoutils.GetMd5(user.Password)
	user.Status = users.StatusActive
	user.DateCreated = dateutils.GetNowDBFormat()
	if err := user.SAVE(); err != nil {
		return nil, err
	}
	return &user, nil

}

// UpdateUser create new user
func (s *usersService) UpdateUser(user users.User) (*users.User, *errors.RestErr) {
	currentUser := &users.User{ID: user.ID}
	if err := currentUser.GET(); err != nil {
		return nil, err
	}
	fmt.Println(currentUser)
	if user.Email != "" {
		currentUser.Email = user.Email
	}
	if user.FirstName != "" {
		currentUser.FirstName = user.FirstName
	}
	if user.LastName != "" {
		fmt.Println("changing LastName")
		currentUser.LastName = user.LastName
	}
	if err := currentUser.UPDATE(); err != nil {
		return nil, err
	}
	fmt.Println(currentUser)
	return currentUser, nil
}

// DeleteUser from database
func (s *usersService) DeleteUser(userID int64) *errors.RestErr {
	deletingUser := &users.User{ID: userID}
	return deletingUser.DELETE()
}

// Search return users by status
func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}

func (s *usersService) LoginUser(request users.LoginRequest) (*users.User, *errors.RestErr) {
	dao := &users.User{
		Email:    request.Email,
		Password: request.Password,
	}
	if err := dao.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return dao, nil
}
