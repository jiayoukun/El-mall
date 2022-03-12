package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"strconv"
	"test15/cache"
)

//商品模型
type Product struct {
	gorm.Model
	ProductId string `gorm:"primary_key;unique"`
	Name string
	Num string `gorm:"default:1"`
	CategoryId string //1为电子类 2为书籍类
	Title string
	ImgPath string
	Price string
	DiscountPrice string
	Info   string `gorm:"size:1000"`
}
//展示点击数
func (product*Product)View() uint64 {
	strs,_:=cache.RedisClient.ZScore("key",product.ProductId).Result()
	s:=fmt.Sprintf("%.f",strs)
	count,_ := strconv.ParseUint(s,10,64)
	fmt.Println("count",count)
	return count
}
//增加当前商品点击数
func (product*Product)AddNowKey()  {
	cache.RedisClient.ZIncrBy("key",1,product.ProductId)
}
//展示每日排行
func (product*Product)TodayRankView()  {
	cache.RedisClient.ZIncrBy(cache.BookKey,1,strconv.Itoa(int(product.ID)))
}
//增加书类点击数
func (product*Product)AddBookRank()  {
	cache.RedisClient.ZIncrBy(cache.BookKey,1,cache.BookKey)
}
//增加电子类点击数
func (product*Product)AddElectronicRank()  {
	fmt.Println("product.ProductId",product.ProductId)
	cache.RedisClient.ZIncrBy(cache.Electronic,1,product.Name)
}
