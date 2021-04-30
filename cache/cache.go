package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"sync"
)
type Cache struct {
	Actual 	map[string]bool
	Items map[string]Item

	sync.RWMutex
}

type LocalCache struct {
	Key		string
	Cache	interface{}
	IsActual bool
}

type Item struct {
	Values interface{}
}

func InitCache() *Cache{
	items := make(map[string]Item)

	cache := Cache{
		Actual: make(map[string]bool),
		Items: items,
	}

	return &cache
}

func (c *Cache) AddToCache(key string, customStruct interface{}) {
	c.Lock()
	defer c.Unlock()

	c.Items[key] = Item{
		Values: customStruct,
	}
	c.Actual[key] = true
}

func (c *Cache) CheckKey(key string) bool {
	_, ok := c.Items[key]
	if !ok {
		return false
	}
	return true
}

func (c *Cache) DeleteCache(key string) {
	c.Lock()
	defer c.Unlock()

	delete(c.Items, key)
}

func (c *Cache) CheckActual(key string) bool{
	return c.Actual[key]
}

func (c *Cache) ChangeActualStatus(key string) {
	c.Actual[key] = false
}

func (c *Cache) GetAllCache() []LocalCache {
	localCache := make([]LocalCache, 0)

	for k, v := range c.Items {
		localCache = append(localCache, LocalCache{
			Key: k,
			Cache:    v,
			IsActual: c.Actual[k],
		})
	}
	return localCache
}
func Decode(data []byte, to interface{}) error {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	return dec.Decode(to)
}


func (c *Cache) GetTestCache(key string, to interface{}) {
	err := Decode(c.Items[key].Values.([]byte), to)
	if err != nil {
		log.Print(err)
	}
}

func (c *Cache) GetCache(key string) interface{} {
	if c.CheckActual(key) {
		c.RLock()
		defer c.RUnlock()
		item, ok := c.Items[key]
		if !ok {
			fmt.Println("Нужно заполнить кэш")
			return nil
		}




		return item.Values
	}

	return nil
}

