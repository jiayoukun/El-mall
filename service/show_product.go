package service

import (
	"github.com/go-redis/redis"
	"test15/cache"
	"test15/model"
	"test15/pkg/e"
	"test15/serializer"
)

type ShowProductService struct {
}

func (service*ShowProductService)Show(id string)serializer.Response  {
	var product model.Product
	var count int
	code:=e.SUCCESS
	model.DB.Model(&product).Where("product_id=?",id).Count(&count)
	if count==0{
		code=e.ERROR_NOT_EXIST_PRODUCT
		return serializer.Response{
			Status:code,
			Msg:e.GetMsg(code),
		}
	}
	model.DB.First(&product,"product_id=?",id)

	if product.CategoryId=="1"{
		_,err:=cache.RedisClient.ZScore("key",product.ProductId).Result()
		if err==nil {
			product.AddNowKey()
		}else {
			cache.RedisClient.ZAdd("key",redis.Z{1,product.ProductId})
		}
		_,err=cache.RedisClient.ZScore(cache.Electronic,product.Name).Result()
		if err==nil {
			product.AddElectronicRank()
		}else {
			cache.RedisClient.ZAdd(cache.Electronic,redis.Z{1,product.Name})
		}
	}else if product.CategoryId=="2"{
		product.AddBookRank()
	}

	return serializer.Response{
		Status: code,
		Msg: e.GetMsg(code),
		Data:serializer.BuildProduct(product),
	}
}
