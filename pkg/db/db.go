package db

import (
	"task-tracker/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Db struct {
	*gorm.DB
}

func NewDb(conf *config.Config) *Db{
	db, err := gorm.Open(postgres.Open(conf.Db.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Db{db}
}
