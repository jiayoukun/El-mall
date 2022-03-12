package serializer

import "test15/model"

type Cart struct {
	ProductID string `json:"product_id"`
	Num string `json:"num"`
	ProductName string `json:"product_name"`
	Imgpath string `json:"imgpath"`
	Price string `json:"price"`
}

func Buildcart(carttemp model.Cart) Cart {
	return Cart{
		carttemp.ProductID,
		carttemp.Num,
		carttemp.ProductName,
		carttemp.Imgpath,
		carttemp.Price,
	}
}

func BuildCartData(carttemp []model.Cart) (temp2 []Cart) {
		for _,items := range carttemp{
			temp:=Buildcart(items)
			temp2=append(temp2,temp)
	}
	return temp2
}