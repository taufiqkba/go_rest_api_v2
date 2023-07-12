package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,unique,min=3,max=200"`
}
