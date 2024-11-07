package requests

import "github.com/google/uuid"

type CreateURLRequest struct {
	OriginalURL string    `json:"original_url" binding:"required,url"`
	UserId      uuid.UUID `json:"user_id" binding:"required,user_id"`
}

type GetUserUrlsRequest struct {
	UserId uuid.UUID `json:"user_id" binding:"required,user_id"`
}

type GetURLRequest struct {
	Id uuid.UUID `json:"id" binding:"required"`
}

type GetUrlByShortNameRequest struct {
	ShortName string `json:"short_name" binding:"required"`
}

type DeleteByShortNameRequest struct {
	ShortName string `json:"short_name" binding:"required"`
}

type GetUrlStatByShortNameRequest struct {
	ShortName string `json:"short_name" binding:"required"`
}

type GetUrlByOriginalNameRequest struct {
	OriginalName string `json:"original_name" binding:"required"`
}
