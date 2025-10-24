package auth

type ResponseStruct struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,max=32"`
}

type RegisterResponse struct {
	ResponseStruct
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type LoginResponse struct {
	ResponseStruct
}
