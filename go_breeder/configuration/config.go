package configuration

import (
	"database/sql"
	"go_breeder/adapters"
	"go_breeder/models"
	"sync"
)

type Application struct {
	Models     *models.Models
	CatService *adapters.RemoteService
}

var instance *Application
var once sync.Once
var db *sql.DB
var CatService *adapters.RemoteService

func New(pool *sql.DB, cs *adapters.RemoteService) *Application {
	db = pool
	CatService = cs
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{Models: models.New(db), CatService: CatService}
	})

	return instance
}
