package redis

import (
	"testing"
)

func TestNewRedis(t *testing.T) {
	redis, err := New(Config{Address: "127.0.0.1:6379"})
	if err != nil {
		t.Fatal(err)
	}
	redis.Close()
}
