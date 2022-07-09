package cache

import "time"

type Cache struct {
	items map[string]Element
}

type Element struct {
	value    string
	deadline time.Time
}

func NewCache() Cache {
	itemsMap := make(map[string]Element)
	c := Cache{
		items: itemsMap,
	}
	return c
}

func (c Cache) Get(key string) (string, bool) {
	if c.items[key].deadline.Before(time.Now()) && !c.items[key].deadline.IsZero() {
		delete(c.items, key)
	}
	el, inCache := c.items[key]
	return el.value, inCache
}

func (c *Cache) Put(key, value string) {
	c.items[key] = Element{
		value:    value,
		deadline: time.Time{},
	}
}

func (c Cache) Keys() []string {
	outputKeys := make([]string, 0, len(c.items))
	for key, elem := range c.items {
		if elem.deadline.Before(time.Now()) && !elem.deadline.IsZero() {
			delete(c.items, key)
		}
		outputKeys = append(outputKeys, key)
	}
	return outputKeys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.items[key] = Element{
		value:    value,
		deadline: deadline,
	}
}
