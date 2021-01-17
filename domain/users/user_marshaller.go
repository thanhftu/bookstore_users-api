package users

import (
	"encoding/json"
)

// PublicUser user from public
type PublicUser struct {
	ID int64 `json:"id"`
	// FirstName   string `json:"first_name"`
	// LastName    string `json:"last_name"`
	// Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`

}

// PrivateUser struct
type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	// Password    string `json:"password"`
}

// Marshall for slice of user
func (users Users) Marshall(isPublic bool) interface{} {
	results := make([]interface{}, len(users))
	for index, user := range users {
		results[index] = user.Marshall(isPublic)
	}
	return results
}

// Marshall return a private or public user
func (user *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err.Error()
	}
	var privateUser PrivateUser
	json.Unmarshal(userJSON, &privateUser)
	return privateUser
}
