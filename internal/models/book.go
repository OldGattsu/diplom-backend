package models

type Book struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Authors     []*Author `json:"authors,omitempty"`
}
