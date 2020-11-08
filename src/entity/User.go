package entity

// User ...
type User struct {
	ID       int    `json:"id"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}
