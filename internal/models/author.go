package models

type Author struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Books []*Book `json:"books,omitempty"`
}
