package lru

import "container/list"

type Value interface {
	Len() int
}

// Strategy different implements of cache strategy
type Strategy interface {
	Add(string, Value)
	Remove(string) Value
	Get(string) (v Value, ok bool)
}

type Cache struct {
	maxBytes  int64
	nBytes    int64
	ll        *list.List // doubly linked list
	cache     map[string]*list.Element
	OnEvicted func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

func NewCache(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:  maxBytes,
		OnEvicted: onEvicted,
		cache:     make(map[string]*list.Element),
		ll:        list.New(),
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if v, ok := c.cache[key]; ok {
		c.ll.MoveToFront(v)
		kv := v.Value.(*entry)
		return kv.value, true
	}
	return
}

// RemoveOldest remove oldest element from lru cache
func (c *Cache) RemoveOldest() {
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}

}

func (c *Cache) Add(key string, value Value) {
	if e, ok := c.cache[key]; ok {
		c.ll.MoveToFront(e)
		kv := e.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		newEle := c.ll.PushFront(&entry{key, value})
		c.cache[key] = newEle
		c.nBytes += int64(len(key)) + int64(value.Len())
	}

	for c.maxBytes != 0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
