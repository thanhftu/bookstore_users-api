package app

import (
	"github.com/thanhftu/bookstore_users-api/controllers/ping"
	"github.com/thanhftu/bookstore_users-api/controllers/users"
)

func mapURL() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:userID", users.GetUser)
	router.PUT("/users/:userID", users.UpdateUser)
	router.DELETE("/users/:userID", users.DeleteUser)
	router.GET("/internal/users/search", users.Search)
}
