package domain

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis"
)

//Broker broker interface
type Broker interface {
	Publish(channel string, message interface{}) *redis.IntCmd
	Subscribe(channel string, cb func(string))
	Set(key string, message interface{}, time time.Duration) error
}

type brokerStruct struct {
	rdb redis.Client
}

var ctx = context.Background()

//NewBroker create a new broker
func NewBroker() Broker {

	redisURL := os.Getenv("REDIS_URL")
	if len(redisURL) == 0 {
		redisURL = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisURL,
		Password: "",
		DB:       0,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		log.Fatal("Broker is down", err)
	}

	return &brokerStruct{
		rdb: *rdb,
	}
}

//Set set key/value
func (b *brokerStruct) Set(key string, message interface{}, time time.Duration) error {
	redisError := b.rdb.Set(ctx, key, message, time).Err()

	if redisError != nil {
		return redisError
	}

	return nil
}

//Publish publish some message in a channel
func (b *brokerStruct) Publish(channel string, message interface{}) *redis.IntCmd {
	return b.rdb.Publish(ctx, channel, message)
}

//Subscribe subscribe to a channel
func (b *brokerStruct) Subscribe(channel string, cb func(string)) {

	go func() {

		psNewMessage := b.rdb.Subscribe(ctx, channel)

		for {

			msg, err := psNewMessage.ReceiveMessage(ctx)

			if err == nil {
				cb(msg.Payload)
			} else {
				fmt.Println(err)
			}
		}

	}()

}
