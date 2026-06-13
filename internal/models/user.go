package models

type UserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=100"`
	Dob  string `json:"dob" validate:"required"`
}

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
}

type UserResponseWithAge struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age,omitempty"`
}
