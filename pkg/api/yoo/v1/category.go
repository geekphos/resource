package v1

import "time"

type CreateCategoryRequest struct {
	Name     string `json:"name"`
	ParentID int32  `json:"parent_id"`
}

type UpdateCategoryRequest struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	ParentID int32  `json:"parent_id"`
}

type AllCategoryResponse struct {
	ID        int32                  `json:"id"`
	Name      string                 `json:"name"`
	ParentID  int32                  `json:"parent_id"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	Children  []*AllCategoryResponse `json:"children"`
}
