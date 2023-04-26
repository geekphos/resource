package v1

import (
	"time"

	"gorm.io/datatypes"
)

type GetResourceResponse struct {
	ID        int32          `json:"id"`
	Name      string         `json:"name"`
	Label     string         `json:"label"`
	Badge     string         `json:"badge"`
	Category  string         `json:"category"`
	Tags      datatypes.JSON `json:"tags"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type ListResourceRequest struct {
	Page     int    `json:"page" form:"page,default=1" binding:"omitempty,gte=1"`
	PageSize int    `json:"page_size" form:"page_size,default=10" binding:"omitempty,gte=1"`
	Name     string `json:"name" form:"name"`
	Label    string `json:"label" form:"label"`
	Category string `json:"category" form:"category"`
	Tag      string `json:"tag" form:"tag"`
}

type ListResourceResponse struct {
	ID        int32          `json:"id"`
	Name      string         `json:"name"`
	Label     string         `json:"label"`
	Badge     string         `json:"badge"`
	Category  string         `json:"category"`
	Tags      datatypes.JSON `json:"tags"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
}

type UpdateResourceRequest struct {
	ID       int32          `json:"id" uri:"id" binding:"required,gte=1"`
	Name     *string        `json:"name"`
	Label    *string        `json:"label"`
	Badge    *string        `json:"badge"`
	Category *string        `json:"category"`
	Tags     datatypes.JSON `json:"tags"`
}

type ListUsedResourceRequest struct {
	Name   string   `json:"name" form:"tag"`
	Tags   []string `json:"tags" form:"tag"`
	Letter string   `json:"letter" form:"letter"`
}

type ListUsedResourceResponse struct {
	ID          int32     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Badge       string    `json:"badge"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
