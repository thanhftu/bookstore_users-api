package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thanhftu/bookstore_oauth-go/oauth"
	"github.com/thanhftu/bookstore_users-api/domain/users"
	"github.com/thanhftu/bookstore_users-api/services"
	"github.com/thanhftu/bookstore_utils-go/resterrors"
)

func getUserID(userid string) (int64, error) {
	return strconv.ParseInt(userid, 10, 64)
}

// CreateUser used when creating new user
func CreateUser(c *gin.Context) {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := resterrors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.UsersService.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result.Marshall(c.GetHeader("x-public") == "true"))
}

// GetUser return users
func GetUser(c *gin.Context) {
	if err := oauth.AuthenticateRequest(c.Request); err != nil {
		c.JSON(err.Status, err)
		return
	}
	userID, err := getUserID(c.Param("userID"))
	if err != nil {
		err := resterrors.NewBadRequestError("user ID should be a number")
		c.JSON(err.Status, err)
		return
	}
	user, getErr := services.UsersService.GetUser(userID)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	if oauth.GetCallerID(c.Request) == user.ID {
		c.JSON(http.StatusOK, user.Marshall(false))
		return
	}
	c.JSON(http.StatusOK, user.Marshall(oauth.IsPublic(c.Request)))
}

// UpdateUser used when creating new user
func UpdateUser(c *gin.Context) {
	userID, userErr := getUserID(c.Param("userID"))
	if userErr != nil {
		err := resterrors.NewBadRequestError("user ID should be a number")
		c.JSON(err.Status, err)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := resterrors.NewBadRequestError("invalid json data")
		c.JSON(restErr.Status, restErr)
		return
	}
	user.ID = userID
	fmt.Println(user)
	result, err := services.UsersService.UpdateUser(user)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, result.Marshall(c.GetHeader("x-public") == "true"))
}

// DeleteUser deletes existing user
func DeleteUser(c *gin.Context) {
	userID, errID := getUserID(c.Param("userID"))
	if errID != nil {
		errID := resterrors.NewBadRequestError("user ID should be a number")
		c.JSON(errID.Status, errID)
		return
	}
	err := services.UsersService.DeleteUser(userID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

// Search users
func Search(c *gin.Context) {
	status := c.Query("status")
	users, err := services.UsersService.SearchUser(status)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	results := users.Marshall(c.GetHeader("x-public") == "true")
	c.JSON(http.StatusOK, results)
}

// Login user
func Login(c *gin.Context) {
	var loginRequest users.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		restErr := resterrors.NewBadRequestError("invalid login json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	user, err := services.UsersService.LoginUser(loginRequest)
	// fmt.Println(loginRequest)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, user)
}
