package order

type orderItem struct {
	ProductId int64 `json:"product_id"`
	Quantity  int64 `json:"quantity"`
}

type placeOrderPrams struct {
	UserId int64       `json:"user_id"`
	Items  []orderItem `json:"items"`
}
