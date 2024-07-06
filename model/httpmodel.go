package model

type PlaceOrderHTTPRequest struct {
	Products []ProductHttp `json:"products" binding:"required,dive"`
}

type ProductHttp struct {
	ProductId string `json:"product_id" binding:"required"`
	Quantity  int32  `json:"quantity" binding:"required,gte=1"`
}

type PlaceOrderResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}
