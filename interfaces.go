package gomockito

import (
	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
)

type (
	SqlMockRepo interface {
		Titles() []string
		Values() []driver.Value
		Row() sqlmock.Rows
		Fresh()
	}
)
