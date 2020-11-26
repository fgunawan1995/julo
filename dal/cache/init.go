package cache

import (
	cache "github.com/patrickmn/go-cache"
)

type DAL struct {
	Cache *cache.Cache
}

func New(cache *cache.Cache) *DAL {
	return &DAL{
		Cache: cache,
	}
}
