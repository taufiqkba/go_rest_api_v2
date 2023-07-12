package web

type CategoryCreateRequest struct {
	Name string `validate:"required,unique,min=3,max=200"`
}
