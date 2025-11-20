package verify

type SendRequestStruct struct {
	Email string `json:"email" validate:"required,email"`
}
