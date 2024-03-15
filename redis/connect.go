package redis

import "github.com/redis/go-redis/v9"

type Redis struct {
	Client       *redis.Client
	Addr 		 string
	Password 	 string
}

func (s *Redis) Connect() {
	s.Client = redis.NewClient(&redis.Options{
        Addr:	  s.Addr,
        Password: s.Password, // no password set
        DB:		  0,  // use default DB
    })
}
