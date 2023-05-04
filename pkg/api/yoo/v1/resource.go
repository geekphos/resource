package v1

import (
	"time"
)

type GetResourceResponse struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Badge       string    `json:"badge"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type ListResourceRequest struct {
	Page        int    `json:"page" form:"page,default=1" binding:"omitempty,gte=1"`
	PageSize    int    `json:"page_size" form:"page_size,default=10" binding:"omitempty,gte=1"`
	Name        string `json:"name" form:"name"`
	Description string `json:"description" form:"description"`
}

type AllResourceRequest struct {
	Description string `json:"description" form:"description"`
}

type ListResourceResponse struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Badge       string    `json:"badge"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateResourceRequest struct {
	ID          int32   `json:"id" uri:"id" binding:"required,gte=1"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Badge       *string `json:"badge"`
}

type ListUsedResourceRequest struct {
	Name string `json:"name" form:"tag"`
}

type ListUsedResourceResponse struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Badge       string    `json:"badge"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
