package dto

type Order struct {
	Id     int    `json:"id"`
	IsPaid bool   `json:"isPaid"`
	Type   string `json:"type"`
}
