package entity

// activity
type ActivityCreateRequest struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type ActivityUpdateRequest struct {
	Title string `json:"title" validate:"required"`
}

// todo
type TodoCreateRequest struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupID uint   `json:"activity_group_id" validate:"required"`
}

type TodoUpdateRequest struct {
	Title    string `json:"title" validate:"required"`
	IsActive bool   `json:"is_active" validate:"omitempty,required"`
}
