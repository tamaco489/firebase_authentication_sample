package repository

import (
	"context"
	"database/sql"
	"fmt"
	"sync"
	"time"

	"github.com/tamaco489/firebase_authentication_sample/api/core/internal/configuration"

	mysql_driver "github.com/go-sql-driver/mysql"
)

type Beginner interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// Beginnerインターフェースがsql.DBのメソッドを定義しているかのチェック
var _ Beginner = (*sql.DB)(nil)

var (
	instance *sql.DB
	once     sync.Once
)

func InitDB() *sql.DB {
	var err error
	once.Do(func() {
		instance, err = connect()
		if err != nil {
			panic(err)
		}
	})
	return instance
}

func connect() (*sql.DB, error) {
	c := mysql_driver.Config{
		User:                 configuration.Get().CoreDB.User,
		Passwd:               configuration.Get().CoreDB.Pass,
		Addr:                 fmt.Sprintf("%s:%s", configuration.Get().CoreDB.Host, configuration.Get().CoreDB.Port),
		DBName:               configuration.Get().CoreDB.Name,
		ParseTime:            true,
		Net:                  "tcp",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return nil, err
	}

	lifetime := 1 * time.Minute
	if configuration.Get().API.Env == "stg" {
		lifetime = 10 * time.Second
	}

	db.SetMaxOpenConns(2)
	db.SetMaxIdleConns(0)
	db.SetConnMaxLifetime(lifetime)

	return db, err
}
