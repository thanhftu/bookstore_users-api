package mysqlutils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/thanhftu/bookstore_users-api/utils/errors"
)

const (
	errNoRowFound = "no rows in result set"
)

// ParseError handles error relating to mysql
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errNoRowFound) {
			return errors.NewNotFoundError(fmt.Sprintf("no user with given id found"))
		}
		return errors.NewInternalServerError("error when parsing database response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError("invalid data")
	}
	return errors.NewInternalServerError("error processing request")
}
