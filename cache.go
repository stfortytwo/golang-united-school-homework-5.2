package cache

import "time"

type Cache struct {
	items map[string]Element
}

type Element struct {
	value string
}

func NewCache() Cache {
	items_map := make(map[string]Element)
	c := Cache{
		items: items_map,
	}
	return c
}

func (c Cache) Get(key string) (string, bool) {
	el, inCache := c.items[key]
	return el.value, inCache
}

func (c *Cache) Put(key, value string) {
	c.items[key] = Element{
		value: value,
	}
}

func (c Cache) Keys() []string {
	outputKeys := make([]string, 0, len(c.items))
	for key := range c.items {
		outputKeys = append(outputKeys, key)
	}
	return outputKeys
}

func (receiver) PutTill(key, value string, deadline time.Time) {
}
