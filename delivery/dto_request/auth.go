package dto_request

type AuthEmailLoginRequest struct {
	Email    string `json:"email" validate:"required,not_empty,email" example:"email@gmail.com"`
	Password string `json:"password" validate:"required,not_empty,alphanum_symbol" example:"123$#25"`
} // @name AuthEmailLoginRequest

type AuthEmailRegisterRequest struct {
	Name     string `json:"name" validate:"required,not_empty" example:"John Doe"`
	Email    string `json:"email" validate:"required,not_empty,email" example:"email@gmail.com"`
	Password string `json:"password" validate:"required,not_empty,alphanum_symbol" example:"123$#25"`
} // @name AuthEmailRegisterRequest

type AuthGoogleLoginRequest struct {
	Code  string `json:"code" validate:"required,not_empty"`
	State string `json:"state" extensions:"x-nullable"`
} // @name AuthGoogleLoginRequest
