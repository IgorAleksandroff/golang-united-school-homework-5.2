package cache

import "time"

type valueWithTime struct {
	value        string
	notCanExpire bool
	deadline     time.Time
}

type Cache struct {
	memo map[string]valueWithTime
}

func NewCache() Cache {
	return Cache{
		make(map[string]valueWithTime),
	}
}

func (c *Cache) Get(key string) (string, bool) {
	valWithTime, ok := c.memo[key]
	return valWithTime.value, ok && (valWithTime.notCanExpire || !time.Now().After(valWithTime.deadline))
}

func (c *Cache) Put(key, value string) {
	c.memo[key] = valueWithTime{
		value:        value,
		notCanExpire: true,
	}
}

func (c *Cache) Keys() []string {
	keys := make([]string, 0, len(c.memo))
	for key := range c.memo {
		if _, ok := c.Get(key); ok {
			keys = append(keys, key)
		}
	}
	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.memo[key] = valueWithTime{
		value:    value,
		deadline: deadline,
	}
}
