package utils

import "container/list"

type entry[K comparable, V any] struct {
	key   K
	value V
}

type LRU[K comparable, V any] struct {
	capacity int
	ll       *list.List
	m        map[K]*list.Element
}

func NewLRU[K comparable, V any](cap int) *LRU[K, V] {
	return &LRU[K, V]{
		capacity: cap,
		ll:       list.New(),
		m:        make(map[K]*list.Element),
	}
}

func (c *LRU[K, V]) Get(k K) (V, bool) {
	if elem, ok := c.m[k]; ok {
		c.ll.MoveToFront(elem)
		return elem.Value.(*entry[K, V]).value, true
	}

	var zero V
	return zero, false
}

func (c *LRU[K, V]) Put(k K, v V) {
	if elem, ok := c.m[k]; ok {
		elem.Value.(*entry[K, V]).value = v
		c.ll.MoveToFront(elem)
		return
	}

	elem := c.ll.PushFront(&entry[K, V]{k, v})
	c.m[k] = elem

	if c.ll.Len() > c.capacity {
		old := c.ll.Back()
		c.ll.Remove(old)
		oldEntry := old.Value.(*entry[K, V])
		delete(c.m, oldEntry.key)
	}
}

func (c *LRU[K, V]) Len() int {
	return c.ll.Len()
}
