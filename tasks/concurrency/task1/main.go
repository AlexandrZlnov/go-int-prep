// собес: авито
// грэйд: midle
// Задача:
// написать кэш ключ-значение с методами Чтение Удаление Добавление
// потокобезопасный
// маштабировать кэш при большом количестве записей

package main

import "sync"

type Shard struct {
	shard       []*Cache
	countShards int
}

type Cache struct {
	mu    sync.RWMutex
	cache map[int]int
}

func NewShardCache(countShards int) *Shard {
	shards := make([]*Cache, countShards)
	for i := range shards {
		shards[i] = &Cache{
			cache: make(map[int]int),
		}
	}
	return &Shard{
		shard:       shards,
		countShards: countShards,
	}
}

func (s *Shard) Set(k int, v int) {
	shardNumber := k % s.countShards
	s.shard[shardNumber].Set(k, v)
}

func (s *Shard) Del(k int) {
	shardNumber := k % s.countShards
	s.shard[shardNumber].Del(k)
}

func (s *Shard) Search(k int) int {
	shardNumber := k % s.countShards
	return s.shard[shardNumber].Search(k)
}

func (c *Cache) Search(k int) int {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.cache[k]
}

func (c *Cache) Set(k int, v int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[k] = v
}

func (c *Cache) Del(k int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.cache, k)
}
