package lrucache

import (
	lru "github.com/hashicorp/golang-lru/v2"
)

var Cache *lru.Cache[string, any]

func InitializeLruCache() {
	Cache, _ = lru.New[string, any](128)
}