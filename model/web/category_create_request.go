package web

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required,min=3,max=200"`
}
