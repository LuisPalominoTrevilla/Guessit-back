package models

// Credentials holds the credentials for login
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AgeGuess holds the guess from a user
type AgeGuess struct {
	Age int `json:"age"`
}
