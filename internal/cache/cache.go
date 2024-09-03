package cache

type MemCache struct {
	cache map[string]interface{}
}

func New() *MemCache {
	cache := make(map[string]interface{})
	return &MemCache{cache}
}

func (mc *MemCache) Get(key string) interface{} {
	return mc.cache[key]
}

func (mc *MemCache) Set(key string, value interface{}) {
	mc.cache[key] = value
}
