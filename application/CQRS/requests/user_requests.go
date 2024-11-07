package requests

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required,email"` // Валидация для корректного email
	Password string `json:"password" binding:"required"`
}

type GetUserRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}

type GetUserByEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}
