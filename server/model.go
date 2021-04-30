package server

import (
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"task4/model"
)

type Server struct {
	api   fasthttp.Server
	log   *zerolog.Logger
	users usersInterface
}

type usersInterface interface {
	GetAllUsers() ([]model.Users, error)
	GetUsersWithMinAge(age int) ([]model.Users, error)
	GetUserById(id int) (model.Users, error)
	CreateUser(user model.Users) error
	UpdateUser(id int, user model.Users) error
	DeleteUser(id int) error
	GetAllCache() interface{}
}
