package models

type GenericResponse struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
	Error   string `json:"error,omitempty"`
}
