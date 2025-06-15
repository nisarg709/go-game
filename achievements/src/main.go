package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"eman/passport/achievements/src/bootstrap"
	"eman/passport/achievements/src/services"
	"encoding/json"
)

func main() {

	service := bootstrap.NewService();
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pubsub := redisClient.Subscribe("events")
	defer pubsub.Close()

	// Wait for confirmation that subscription is created before publishing anything.
	_, err := pubsub.Receive()
	if err != nil {
		panic(err)
	}

	for msg := range pubsub.Channel() {

		event := services.Event{}
		err := json.Unmarshal([]byte(msg.Payload), &event)

		if err != nil {
			fmt.Println("Invalid event format")
			continue
		}

		if err := service.Work(&event); err != nil {
			fmt.Printf("Event %s encountered an error: %e \n", event.Action, err.Error())
		} else {
			fmt.Printf("Event %s handled successfully \n", event.Action)
		}
	}
}
