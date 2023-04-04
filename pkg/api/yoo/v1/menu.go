package v1

type CreateMenuRequest struct {
	Name       string `json:"name" binding:"required"`
	MenuType   int32  `json:"menu_type" binding:"required"`
	ResourceID int32  `json:"resource_id"`
	Href       string `json:"href"`
	ParentID   int32  `json:"parent_id"`
	Number     int32  `json:"number"`
}

type UpdateMenuRequest struct {
	ID         int32   `json:"id" uri:"id" binding:"required"`
	Name       *string `json:"name"`
	MenuType   *int32  `json:"menu_type"`
	ResourceID *int32  `json:"resource_id"`
	Href       *string `json:"href"`
	ParentID   *int32  `json:"parent_id"`
	Number     *int32  `json:"number"`
}

type ListMenuRequest struct {
	Name     string `json:"name"`
	ParentID int32  `json:"parent_id" binding:"omitempty,gt=0"`
}

type ListMenuResponse struct {
	ID         int32               `json:"id"`
	Name       string              `json:"name" binding:"required"`
	MenuType   int32               `json:"menu_type" binding:"required"`
	ResourceID int32               `json:"resource_id"`
	Href       string              `json:"href"`
	ParentID   int32               `json:"parent_id"`
	Number     int32               `json:"number"`
	Children   []*ListMenuResponse `json:"children"`
}

type GetMenuRequest struct {
	Name     string `json:"name"`
	ParentID int32  `json:"parent_id"`
}

type GetMenuResponse struct {
	ID         int32  `json:"id"`
	Name       string `json:"name" binding:"required"`
	MenuType   int32  `json:"menu_type" binding:"required"`
	ResourceID int32  `json:"resource_id"`
	Href       string `json:"href"`
	ParentID   int32  `json:"parent_id"`
	Number     int32  `json:"number"`
}
