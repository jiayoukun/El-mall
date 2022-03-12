package service

import (
	"github.com/go-redis/redis"
	"test15/cache"
	"test15/pkg/e"
	"test15/serializer"
)

type ShowElectronicRankService struct {
}

func (servicr*ShowElectronicRankService)Show() serializer.Response {
	code:=e.SUCCESS
	op := redis.ZRangeBy{
		Min:"0", // 最小分数
		Max:"999999", // 最大分数
		Offset:0, // 类似sql的limit, 表示开始偏移量
		Count:10, // 一次返回多少数据
	}
	result,err:=cache.RedisClient.ZRevRangeByScoreWithScores(cache.Electronic,op).Result()
	if err!=nil{
		code=e.ERROR_DATABASE
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data:result,
	}
}
