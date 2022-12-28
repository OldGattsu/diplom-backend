package models

type Book struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Authors []*Author `json:"authors,omitempty"`
}
