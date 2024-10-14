package payload

type CreateAnimalRequest struct {
	Name  string `json:"name" form:"name" validate:"required"`
	Class string `json:"class" form:"class" validate:"required"`
	Legs  uint   `json:"legs" form:"legs" validate:"required,numeric"`
}

type UpdateAnimalRequest struct {
	Name  string `json:"name" form:"name"`
	Class string `json:"class" form:"class"`
	Legs  uint   `json:"legs" form:"legs" validate:"numeric"`
}

type FindAnimalResponse struct {
	ID    uint   `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Class string `json:"class" form:"class"`
	Legs  uint   `json:"legs" form:"legs"`
}
