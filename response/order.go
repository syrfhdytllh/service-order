package response

type GetAllData struct {
	Customer_name string     `json:"customerName"`
	Ordered_at    string     `json:"orderAt"`
	Items         []ItemShow `json:"items"`
}

type ItemShow struct {
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}

type OrderCreate struct {
	Customer_name string       `json:"customer_name"`
	Ordered_at    string       `json:"ordered_at"`
	Items         []ItemCreate `json:"items"`
}

type ItemCreate struct {
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    string `json:"quantity"`
}
