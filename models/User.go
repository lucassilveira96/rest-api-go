package userModel

type User struct {
	Id    int    `json:"id"`
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Email string `json:"email"`
	Senha string `json:"senha"`
}

var Pessoas []User
