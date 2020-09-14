package domain

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

//Broker broker interface
type Broker interface {
	Publish(channel string, message interface{}) *redis.IntCmd
	Subscribe(channel string, cb func(interface{}))
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

	pong, err := rdb.Ping(ctx).Result()

	fmt.Print(pong, err)

	//subscribe(*rdb)

	return &brokerStruct{
		rdb: *rdb,
	}
}

//Publish publish some message in a channel
func (b *brokerStruct) Publish(channel string, message interface{}) *redis.IntCmd {
	return b.rdb.Publish(ctx, channel, message)
}

//Subscribe subscribe to a channel
func (b *brokerStruct) Subscribe(channel string, cb func(interface{})) {
	go func() {

		psNewMessage := b.rdb.Subscribe(ctx, channel)
		for {
			msg, _ := psNewMessage.ReceiveMessage(ctx)
			cb(msg)
		}

	}()

}
