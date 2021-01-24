package users

import (
	"fmt"

	"github.com/thanhftu/bookstore_users-api/datasources/mysql/usersdb"
	"github.com/thanhftu/bookstore_users-api/utils/errors"
	"github.com/thanhftu/bookstore_utils-go/logger"
)

const (
	queryInsertUser             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?,?,?);"
	queryGetUser                = "SELECT id, first_name, last_name, email,date_created, status FROM users WHERE id=?;"
	queryUpdateUser             = "UPDATE users SET first_name=?, last_name=?, email=?, status=? WHERE id=?;"
	queryDeleteUser             = "DELETE FROM users WHERE id=?;"
	queryFindUserByStatus       = "SELECT id, first_name, last_name, email,date_created, status FROM users WHERE status=?;"
	queryFindByEmailAndPassword = "SELECT id, first_name, last_name, email,date_created, status FROM users WHERE email=? AND password=?;"
)

// GET get a user in database
func (user *User) GET() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user stmt", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.ID)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to Query get user from database", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

// SAVE a user into database
func (user *User) SAVE() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare saving stmt", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	insertResult, saveErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if saveErr != nil {
		logger.Error("error when trying to executing query saving user into database", saveErr)
		return errors.NewInternalServerError("database error")
	}
	userID, errInsrt := insertResult.LastInsertId()
	if errInsrt != nil {
		logger.Error("error when getting last Insert ID", errInsrt)
		return errors.NewInternalServerError("database error")
	}
	user.ID = userID
	return nil

}

// UPDATE updates user
func (user *User) UPDATE() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when preparing update query", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.ID)
	if errExec != nil {
		logger.Error("error when executing update user from database", errExec)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

// DELETE delete user
func (user *User) DELETE() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when preaparing Delete query", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, errExec := stmt.Exec(user.ID)
	if errExec != nil {
		logger.Error("error when executing delete user from database", errExec)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

// FindByStatus return users with given status
func (user *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := usersdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when preparing find user query database", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when excuting stmt finding user from db", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when scan user row into user struct", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no user with status %s found", status))
	}
	return results, nil
}

// FindByEmailAndPassword get a user from database by matching email and password
func (user *User) FindByEmailAndPassword() *errors.RestErr {
	stmt, err := usersdb.Client.Prepare(queryFindByEmailAndPassword)
	if err != nil {
		logger.Error("error when trying to prepare finding user stmt by using email and password", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Email, user.Password)
	if err := result.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
		logger.Error("error when trying to Query get user from database by email and password", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}
