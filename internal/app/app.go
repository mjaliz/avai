package app

import "gorm.io/gorm"

var App Application

type Application struct {
	db *gorm.DB
}

func NewApp(db *gorm.DB) {
	App = Application{db: db}
}

func (a Application) DB() *gorm.DB {
	return a.db
}
