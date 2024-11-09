package responses

import "github.com/google/uuid"

type GetUserResponse struct {
	Id       *uuid.UUID `json:"id"`
	Username *string    `json:"username"`
	Email    *string    `json:"email"`
	Error    error      `json:"error"`
}

type CreateUserResponse struct {
	Id       *uuid.UUID `json:"id"`
	Username *string    `json:"username"`
	Email    *string    `json:"email"`
	Error    error      `json:"error"`
}

type GetUserByEmailResponse struct {
	Id       *uuid.UUID `json:"id"`
	Username *string    `json:"username"`
	Email    *string    `json:"email"`
	Error    error      `json:"error"`
}
