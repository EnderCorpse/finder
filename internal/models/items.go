package models

type ItemObject struct {
	ID          string `json:"id"`
	Location    string `json:"location"`
	Quantity    int    `json:"quantity"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Requester   string `json:"requester"`
	RequestDate string `json:"request_date"`
	ReturnDate  string `json:"return_date"`
}
