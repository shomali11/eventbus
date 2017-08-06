package cmap

import (
	"github.com/shomali11/util/hashes"
)

const (
	defaultShards = 16
)

type ShardedConcurrentMapOption func(*ShardedConcurrentMap)

func WithNumberOfShards(numberOfShards uint32) ShardedConcurrentMapOption {
	return func(shardedConcurrentMap *ShardedConcurrentMap) {
		if numberOfShards < 1 {
			shardedConcurrentMap.numberOfShards = defaultShards
		} else {
			shardedConcurrentMap.numberOfShards = numberOfShards
		}
	}
}

// NewShardedConcurrentMap creates a new sharded concurrent map
func NewShardedConcurrentMap(options ...ShardedConcurrentMapOption) *ShardedConcurrentMap {
	shardedConcurrentMap := &ShardedConcurrentMap{
		numberOfShards: defaultShards,
	}

	for _, option := range options {
		option(shardedConcurrentMap)
	}

	internalMap := make([]*ConcurrentMap, shardedConcurrentMap.numberOfShards)
	for i := uint32(0); i < shardedConcurrentMap.numberOfShards; i++ {
		internalMap[i] = NewConcurrentMap()
	}

	shardedConcurrentMap.internalMap = internalMap
	return shardedConcurrentMap
}

// ShardedConcurrentMap concurrent map
type ShardedConcurrentMap struct {
	numberOfShards uint32
	internalMap    []*ConcurrentMap
}

// Set concurrent set to map
func (c *ShardedConcurrentMap) Set(key string, value interface{}) {
	shard := c.getShard(key)
	concurrentMap := c.internalMap[shard]
	concurrentMap.Set(key, value)
}

// Get concurrent get from map
func (c *ShardedConcurrentMap) Get(key string) (interface{}, bool) {
	shard := c.getShard(key)
	concurrentMap := c.internalMap[shard]
	return concurrentMap.Get(key)
}

// Remove concurrent remove from map
func (c *ShardedConcurrentMap) Remove(key string) {
	shard := c.getShard(key)
	concurrentMap := c.internalMap[shard]
	concurrentMap.Remove(key)
}

func (c *ShardedConcurrentMap) getShard(key string) uint32 {
	return hashes.FNV32(key) % uint32(c.numberOfShards)
}
