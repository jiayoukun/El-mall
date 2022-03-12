package serializer

import "test15/model"

type Order struct {
	ProductID string `json:"product_id"`
	SumPrice string `json:"sum_price"`
	Num string `json:"num"`
	OrderNum  string `json:"order_num"`
	Address string `json:"address"`
	AddressName string `json:"address_name"`
	AddressPhone string `json:"address_phone"`
}

func Buildorder(temp model.Order) Order {
	return Order{
		temp.ProductID,
		temp.SumPrice,
		temp.Num,
		temp.OrderNum,
		temp.Address,
		temp.AddressName,
		temp.AddressPhone,
	}
}

func BuildOrderData(temp []model.Order) (temp3 []Order) {
	for _,temp1 := range temp{
		temp2 := Buildorder(temp1)
		temp3 = append(temp3,temp2)
	}
	return temp3
}