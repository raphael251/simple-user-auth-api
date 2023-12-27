package dto

type CreateUserInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
