package mysqlutils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/thanhftu/bookstore_utils-go/resterrors"
)

const (
	// ErrNoRowFound no row found in db
	ErrNoRowFound = "no rows in result set"
)

// ParseError handles error relating to mysql
func ParseError(err error) *resterrors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrNoRowFound) {
			return resterrors.NewNotFoundError(fmt.Sprintf("no user with given id found"))
		}
		return resterrors.NewInternalServerError("error when parsing database response", errors.New("database error"))
	}
	switch sqlErr.Number {
	case 1062:
		return resterrors.NewBadRequestError("invalid data, email already existed")
	}
	return resterrors.NewInternalServerError("error processing request", errors.New("database error"))
}
