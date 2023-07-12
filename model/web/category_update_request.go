package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `json:"name" validate:"required,min=3,max=200"`
}
