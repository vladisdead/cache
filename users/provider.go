package users

import (
	"cache"
	"github.com/rs/zerolog"
)

func NewProvider(
	storageProvider storageInterface, log *zerolog.Logger, cache *cache.Cache) *Provider {
	p := Provider{
		storage: storageProvider,
		log:     log,
		cache: 	cache,
	}

	return &p
}
