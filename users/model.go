package users

import (
	"cache"
	"github.com/rs/zerolog"
	"task4/model"
)

type Provider struct {
	storage storageInterface
	log     *zerolog.Logger
	cache	*cache.Cache
}

type storageInterface interface {
	GetAllUsersFromDB() ([]model.Users, error)
	GetUsersWithMinAgeFromDB(age int) ([]model.Users, error)
	GetUserByIdDB(id int) (model.Users, error)
	CreateUserDB(user model.Users) error
	ChangeUserDB(id int, user model.Users) error
	DeleteUserDB(id int) error
	GetAllCache() ([]model.LocalCache,error)
}
