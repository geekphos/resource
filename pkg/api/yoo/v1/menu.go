package v1

import "gorm.io/datatypes"

type CreateMenuRequest struct {
	Name        string         `json:"name" binding:"required"`
	MenuType    int32          `json:"menu_type" binding:"required"`
	Icon        string         `json:"icon"`
	ResourceID  *int32         `json:"resource_id"`
	Categories  datatypes.JSON `json:"categories"`
	Description string         `json:"description"`
	Href        *string        `json:"href"`
	ParentID    *int32         `json:"parent_id"`
	Number      int32          `json:"number"`
}

type UpdateMenuRequest struct {
	ID          int32          `json:"id" uri:"id" binding:"required"`
	Name        *string        `json:"name"`
	Icon        *string        `json:"icon"`
	MenuType    *int32         `json:"menu_type"`
	ResourceID  *int32         `json:"resource_id"`
	Categories  datatypes.JSON `json:"categories"`
	Description *string        `json:"description"`
	Href        *string        `json:"href"`
	ParentID    *int32         `json:"parent_id"`
	Number      *int32         `json:"number"`
}

type ListMenuRequest struct {
	Name       string   `json:"name" form:"name"`
	Letter     string   `json:"letter" form:"letter"`
	Categories []string `json:"categories" form:"categories"`
}

type ListMenuResponse struct {
	ID          int32               `json:"id"`
	Name        string              `json:"name"`
	Letter      string              `json:"letter"`
	Icon        string              `json:"icon"`
	MenuType    int32               `json:"menu_type"`
	ResourceID  int32               `json:"resource_id"`
	Categories  datatypes.JSON      `json:"categories"`
	Description string              `json:"description"`
	Href        string              `json:"href"`
	ParentID    int32               `json:"parent_id"`
	Number      int32               `json:"number"`
	Depth       int32               `json:"depth"`
	Children    []*ListMenuResponse `json:"children"`
}

type GetMenuRequest struct {
	Name     string `json:"name"`
	Letter   string `json:"letter"`
	ParentID int32  `json:"parent_id"`
}

type GetMenuResponse struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Letter      string         `json:"letter"`
	Icon        string         `json:"icon"`
	MenuType    int32          `json:"menu_type"`
	ResourceID  int32          `json:"resource_id"`
	Categories  datatypes.JSON `json:"categories"`
	Description *string        `json:"description"`
	Href        string         `json:"href"`
	ParentID    int32          `json:"parent_id"`
	Number      int32          `json:"number"`
}

type GetMenuBranchResponse struct {
	ID          int32                    `json:"id"`
	Name        string                   `json:"name"`
	Letter      string                   `json:"letter"`
	Icon        string                   `json:"icon"`
	MenuType    int32                    `json:"menu_type"`
	ResourceID  int32                    `json:"resource_id"`
	Categories  datatypes.JSON           `json:"categories"`
	Description string                   `json:"description"`
	Href        string                   `json:"href"`
	ParentID    int32                    `json:"parent_id"`
	Number      int32                    `json:"number"`
	Depth       int32                    `json:"depth"`
	Children    []*GetMenuBranchResponse `json:"children"`
}

type GetLeaveMenuRequest struct {
	Name       string   `json:"name"`
	Letter     string   `json:"letter"`
	Categories []string `json:"categories"`
}

type GetLeaveMenuResponse struct {
	ID          int32               `json:"id"`
	Name        string              `json:"name"`
	Letter      string              `json:"letter"`
	Icon        string              `json:"icon"`
	MenuType    int32               `json:"menu_type"`
	ResourceID  int32               `json:"resource_id"`
	Categories  datatypes.JSON      `json:"categories"`
	Description string              `json:"description"`
	Href        string              `json:"href"`
	ParentID    int32               `json:"parent_id"`
	Number      int32               `json:"number"`
	Depth       int32               `json:"depth"`
	Children    []*ListMenuResponse `json:"children"`
}
