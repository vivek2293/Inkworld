package store

// Book represents a book in the store. - Used for queries and getting response from db
type Book struct {
	ID     string
	Title  string
	Genre  string
	Author string
	Price  string
}

// Queries
