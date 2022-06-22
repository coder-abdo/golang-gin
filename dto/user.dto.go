package dto

type UserUpdateDto struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6"`
}

type UserCreate struct {
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"email" validate:"email"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"`
}
