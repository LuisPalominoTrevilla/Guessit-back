package models

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	LastName string `json:"lastName"`
	Password string `json:"password"`
	Age      int    `json:"age"`
}
