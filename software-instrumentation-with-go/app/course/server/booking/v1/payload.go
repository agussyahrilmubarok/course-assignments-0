package v1

type CreateBookingRequest struct {
	BatchCode    string `json:"batch_code" binding:"required"`
	CustomerName string `json:"customer_name" binding:"required"`
}
