package contracts

type RealmCreateRequest struct {
	Name string `json:"name"  binding:"required,omitempty,min=3,max=30" validate:"required,omitempty,min=3,max=30"`
}
