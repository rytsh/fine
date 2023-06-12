package msg

type WebApiError struct {
	Err string `json:"error,omitempty" example:"problem explanation"`
}

type WebApiSuccess struct {
	Msg string `json:"msg,omitempty" example:"success"`
}
