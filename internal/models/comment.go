package models

type Comment struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	BookID int    `json:"book_id"`
	Text   string `json:"text"`
	User   *User  `json:"user"`
	Book   *Book  `json:"book"`
}
