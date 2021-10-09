package dto

type CreateUserDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type EditUserDTO struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
