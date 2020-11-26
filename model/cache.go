package model

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/patrickmn/go-cache"
)

//Cache will save user tokens
var Cache *cache.Cache

const (
	cacheFile       = "cache.json"
	cacheExpiration = 24 * time.Hour
	cacheCleanup    = 5 * time.Minute
	SecretKey       = "abcde"
)

func GetCache() *cache.Cache {
	return Cache
}

//InitCache create new cache
func InitCache() {
	var savedCache map[string]cache.Item
	var err error
	defer func() {
		if err != nil {
			log.Println(err)
		}
		Cache = cache.NewFrom(cacheExpiration, cacheCleanup, savedCache)
	}()
	jsonFile, err := os.Open(cacheFile)
	if err != nil {
		return
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(byteValue), &savedCache)
	if err != nil {
		return
	}
	return
}

func SaveCacheToFile() error {
	items := Cache.Items()
	file, err := json.Marshal(items)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(cacheFile, file, 0777)
	if err != nil {
		return err
	}
	return nil
}
