package mysql

import (
	"github.com/go-sql-driver/mysql"
	error_helper "github.com/pgrau/bookstore-user-api/lib/error"
	"strings"
)

const(
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *error_helper.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return error_helper.NotFound("No record matching given id")
		}
		return error_helper.NotFound("No record matching given id")
	}

	switch sqlErr.Number{
	case 1062:
		return error_helper.Conflict("invalid data")
	}

	return error_helper.InternalServerError("error processing request")
}
