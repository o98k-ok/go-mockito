package sqlmocked

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Nested struct {
	Job string `gorm:"column:job"`
}

type TestStruct struct {
	Name  string `gorm:"column:name"`
	Empty string `json:"okok"`
	Age   string `gorm:"column:age"`
	Nested
}

func realDBConn() *sql.DB {
	db, err := MysqlInitDB("root:mysql123@tcp(127.0.0.1:3306)/bill?parseTime=true&charset=utf8mb4").DB()
	if err != nil {
		panic("")
	}
	return db
}

func TestNewUnitTestDBInstance(t *testing.T) {
	defer NewUnitTestDBInstance(nil, realDBConn, false).Close()
	defer NewUnitTestDBInstance(func(conn *gorm.DB, mocker sqlmock.Sqlmock) {
		// setting your conn
		// xxx = conn
		// mocker your data
		// mocker.Expect
	}, nil, true).Close()
}
