package pokecache

import (
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	interval := 5 * time.Second
	cache := NewCache(interval)
	cases := []struct {
		key         string
		value       []byte
		wait        time.Duration
		shouldExist bool
	}{
		{
			key:         "hola",
			value:       []byte("Hola jaja"),
			wait:        1 * time.Second,
			shouldExist: true,
		},
		{
			key:         "muy tarde",
			value:       []byte("has tardado mucho"),
			wait:        5 * time.Second,
			shouldExist: false,
		},
		{
			key:         "a tiempo",
			value:       []byte("justo a tiempo"),
			wait:        500 * time.Millisecond,
			shouldExist: true,
		},
	}
	for _, c := range cases {
		cache.Add(c.key, c.value)
		time.Sleep(c.wait)
		_, ok := cache.Get(c.key)
		if ok != c.shouldExist {
			t.Errorf("the '%v' key should be: %v - found %v", c.key, c.shouldExist, ok)
		}
	}
}
