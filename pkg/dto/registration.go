package dto

type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type RegistrationRequest struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type RegistrationResponse struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
