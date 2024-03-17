package redis

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client   *redis.Client
	Addr     string
	Password string
}

func (s *Redis) Connect() {
	s.Client = redis.NewClient(&redis.Options{
		Addr:     s.Addr,
		Password: s.Password, // no password set
		DB:       0,          // use default DB
	})

	if s.Client == nil {
		log.Printf("connect redis wrong")
		return
	}

	if status, err := s.Client.Ping(context.Background()).Result(); err != nil{
		log.Print(status)
		log.Printf("connect redis wrong")
		return
	}
	log.Printf("connect redis success")
}
