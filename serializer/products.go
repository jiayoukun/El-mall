package serializer

import "test15/model"

type Product struct {
	ProductId	string `json:"product_id"`
	Name string `json:"name"`
	Prince string `json:"prince"`
	DiscountPrice string `json:"discount_price"`
	Imgpath string `json:"imgpath"`
	Title string `json:"title"`
	CategoryId string `json:"category_id"`
	Info string `json:"info"`
	View uint64 `json:"view"`
	CreatedAt     int64  `json:"created_at"`
}
type TEMP struct {
	ProductId	string `json:"product_id"`
	Name string `json:"name"`
	Prince string `json:"prince"`
	DiscountPrice string `json:"discount_price"`
	Imgpath string `json:"imgpath"`
	Title string `json:"title"`
	CategoryId string `json:"category_id"`
}
//序列化商品
func BuildProduct(item model.Product) Product {
	return Product{
		ProductId: item.ProductId,
		Name: item.Name,
		Prince: item.Price,
		DiscountPrice: item.DiscountPrice,
		Imgpath: item.ImgPath,
		Title: item.Title,
		CategoryId: item.CategoryId,
		Info: item.Info,
		View:	item.View(),
		CreatedAt: item.CreatedAt.Unix(),
	}
}

func Buildproduct1(item model.Product) TEMP {
	return TEMP{
		item.ProductId,
		item.Name,
		item.Price,
		item.DiscountPrice,
		item.ImgPath,
		item.Title,
		item.CategoryId,
	}
}

func BuildProductData(temp []model.Product) (temp3 []Product) {
	for _,temp1 := range temp {
		temp2 := BuildProduct(temp1)
		temp3 = append(temp3,temp2)
	}
	return temp3
}
