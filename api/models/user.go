package models

type User struct {
	FullName string `json:"full_name"`
	Email    string `json:"user_name"`
	Password string `json:"password"`
}

type Brand struct {
	Name string `json:"name"`
}
