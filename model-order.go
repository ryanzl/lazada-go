package lazada

type GetMultiOrderItemsResponse struct {
	Code                 string   `json:"code,omitempty"`
	Data                 []*GetMultiOrderItemsResponseData  `json:"data,omitempty"`
	RequestId            string   `json:"request_id,omitempty"`
}
type GetMultiOrderItemsResponseData struct {
	OrderNumber          int64       `json:"order_number,omitempty"`
	OrderId              int64       `json:"order_id,omitempty"`
	OrderItems           []*OrderItem `json:"order_items,omitempty"`
}
