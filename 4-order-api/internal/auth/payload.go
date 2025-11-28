package auth

type LoginRequest struct {
	Phone string `json:"phone"`
}

type LoginResponse struct {
	SessionID string `json:"sessionId"`
}

type LoginConfirmRequest struct {
	SessionID string `json:"sessionId"`
	Code      string `json:"code"`
}

type LoginConfirmResponse struct {
	Token string `json:"token"`
}
