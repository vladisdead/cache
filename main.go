package main

import (
	c "cache"

	"github.com/rs/zerolog"
	"task4/cfg"
	"task4/server"
	"task4/storage"
	"task4/users"
)

func main() {
	log := server.NewLogger()

	CFG, err := cfg.New()
	if err != nil {
		log.Err(err)
	}
	InitApi(CFG, log)
}

func InitApi(cfg *cfg.CFG, log *zerolog.Logger) {
	storageProvider, err := storage.NewProvider(cfg.Storage.Connstring, log)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка при инициализации storage")
	}

	oldCache := c.InitCache()



	usersProvider := users.NewProvider(storageProvider, log, oldCache)

	api := server.New(log, usersProvider)

	c.InitDB()
	api.Start()
}
