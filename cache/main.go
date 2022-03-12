package cache

import (
	"github.com/go-redis/redis"
	logging "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

var RedisClient *redis.Client
func Redis()  {
	db,_ := strconv.ParseUint(os.Getenv("REDIS_DB"),10,64)
	client :=redis.NewClient(&redis.Options{
	Addr: os.Getenv("REDIS_ADDR"),
	Password: os.Getenv("REDIS_PW"),
	DB:int(db),
	})
	_,err := client.Ping().Result()
	if err != nil {
		logging.Info(err)
		panic(err)
	}
	RedisClient = client
}
