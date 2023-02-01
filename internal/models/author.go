package models

type Author struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Books       []*Book `json:"books,omitempty"`
}
