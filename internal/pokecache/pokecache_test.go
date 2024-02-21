package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	c := NewCache(time.Minute * 5)
	if c.cache == nil {
		t.Error("cache is nil")
	}
}

func TestSetGet(t *testing.T) {
	cache := NewCache(time.Minute * 5)

	cases := []struct {
		key   string
		value []byte
	}{
		{
			key:   "key1",
			value: []byte("val1"),
		},
		{
			key:   "",
			value: []byte("val2"),
		},
		{
			key:   "key3",
			value: []byte(""),
		},
	}

	for _, c := range cases {
		cache.Set(c.key, c.value)
		actual, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("%v not found", c.key)
		}

		if string(actual) != string(c.value) {
			t.Errorf("Value doesn't match: %v, %v", string(actual), string(c.value))
		}
	}
}

func TestReap(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Set("key1", []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get("key1")

	if ok {
		t.Errorf("%s should have been reaped", "key1")
	}
}

func TestReapFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	cache.Set("key1", []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get("key1")

	if !ok {
		t.Errorf("%s should not have been reaped", "key1")
	}
}
