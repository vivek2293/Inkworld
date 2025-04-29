package models

// BookDetailsResponse represents the response structure for book details.
type BookDetailsResponse struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
	Price  string `json:"price"`
}
