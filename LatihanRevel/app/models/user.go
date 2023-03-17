package models

type User struct {
	ID       int    `json:"ID"`
	Age      int    `json:"Age"`
	Addres   string `json:"Addres"`
	Name     string `json:"Name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType int    `json:"userType"`
}
type UserResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}
