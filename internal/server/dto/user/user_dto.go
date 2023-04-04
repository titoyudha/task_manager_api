package server

type RegisterUserDTO struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

type LogInUserDTO struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"_"`
}

type UpdateUserDTO struct {
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

type GetUserInfoByIdDTO struct {
	ID int64 `json:"id,omitempty"`
}

type GetAllUserDTO struct {
	ID int64 `json:"id,omitempty"`
}
