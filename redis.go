package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Address  string
	Password string
	DB       int
}

type Redis struct {
	config Config
	client *redis.Client
}

func New(config Config) (*Redis, error) {
	r := create(config)
	err := r.connect()
	return &r, err
}
func Create(config Config) (Redis, error) {
	r := create(config)
	err := r.connect()
	return r, err
}
func create(config Config) Redis {
	return Redis{
		config: config,
	}
}
func (r *Redis) connect() error {
	r.client = redis.NewClient(&redis.Options{
		Addr:     r.config.Address,
		Password: r.config.Password,
		DB:       r.config.DB,
	})
	_, err := r.client.Ping(context.Background()).Result()
	return err
}

func (r *Redis) Close() error {
	return r.client.Close()
}

func (r *Redis) Scan(ctx context.Context, pattern string, count int64) ([]string, error) {
	var cursor uint64
	var ret []string
	for {
		keys, cursor, err := r.client.Scan(ctx, cursor, pattern, count).Result()
		if err != nil {
			return keys, err
		}
		ret = append(ret, keys...)
		if cursor == 0 {
			break
		}
	}
	return ret, nil

}

func (r *Redis) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	return r.client.Subscribe(ctx, channel)
}

func (r *Redis) Publish(ctx context.Context, channel string, message interface{}) error {
	return r.client.Publish(ctx, channel, message).Err()
}

func (r *Redis) AddKey(ctx context.Context, key string) error {
	return r.client.Set(ctx, key, "", 0).Err()
}

func (r *Redis) DelKey(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}

type PubSub = redis.PubSub
