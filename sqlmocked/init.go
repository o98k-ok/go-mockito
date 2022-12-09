package sqlmocked

import (
	"database/sql"
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GlobalGormConfig you can custom config yourself
var GlobalGormConfig = &gorm.Config{
	SkipDefaultTransaction: true,
}

type DBInitFn func() *sql.DB
type SqlMockInitFn func(conn *gorm.DB, mocker sqlmock.Sqlmock)

// NewUnitTestDBInstance mock conn by gorm via sqlmock or init src db ny dbInit
func NewUnitTestDBInstance(mockInit SqlMockInitFn, dbInit DBInitFn, localMock bool) *sql.DB {
	if mockInit != nil && localMock {
		db, mocker, _ := sqlmock.New()
		msql := mysql.New(mysql.Config{
			SkipInitializeWithVersion: true,
			Conn:                      db,
		})
		dbConn, err := gorm.Open(msql, GlobalGormConfig)
		if err != nil {
			panic(fmt.Errorf("gorm open %s failed %v", "sqlmock", err))
		}

		mockInit(dbConn, mocker)
		return db
	} else {
		return dbInit()
	}
}

// MysqlInitDB by mysqlDSN
func MysqlInitDB(mysqlDSN string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(mysqlDSN), GlobalGormConfig)

	if err != nil {
		panic(fmt.Errorf("gorm open %s failed %v", mysqlDSN, err))
	}
	return db.Session(&gorm.Session{})
}
