package app

import "gorm.io/gorm"

var App application

type application struct {
	db *gorm.DB
}

func NewApp(db *gorm.DB) {
	App = application{db: db}
}

func (a application) DB() *gorm.DB {
	return a.db
}
