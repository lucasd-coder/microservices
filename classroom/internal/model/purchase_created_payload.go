package model

type Customer struct {
	AuthUserId string `json:"authUserId"`
}

type Product struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Slug  string `json:"slug"`
}

type PurchaseCreatedPayload struct {
	Customer Customer `json:"customer"`
	Product  Product  `json:"product"`
}
