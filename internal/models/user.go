package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	IsBlocked bool   `json:"is_blocked"`
	IsAdmin   bool   `json:"is_admin"`
}
