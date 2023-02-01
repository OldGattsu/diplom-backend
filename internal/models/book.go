package models

type Book struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Poster      string    `json:"poster"`
	Authors     []*Author `json:"authors,omitempty"`
}
