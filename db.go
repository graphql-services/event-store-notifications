package main

import (
	"net/url"
	"strings"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	db *gorm.DB
}

func NewDB(db *gorm.DB) *DB {
	v := DB{db}
	return &v
}

func NewDBWithString(urlString string) *DB {
	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	urlString = strings.Replace(urlString, u.Scheme+"://", "", 1)

	db, err := gorm.Open(u.Scheme, urlString)
	if err != nil {
		panic(err)
	}
	return NewDB(db)
}

func (db *DB) AutoMigrate(values ...interface{}) {
	db.db.AutoMigrate(values...)
}
func (db *DB) Close() error {
	return db.db.Close()
}
