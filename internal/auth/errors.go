package auth

import (
	"errors"

	"github.com/go-sql-driver/mysql"
)

func HandleDBError(err error) *mysql.MySQLError {
	var mysqlErr *mysql.MySQLError

	if errors.As(err, &mysqlErr) {
		return mysqlErr
	} else {
		return nil
	}
}
