package models

// Generic response for success/failure
type GenericResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type HcResponse struct {
	Status string `json:"status"`
	Message string `json:"hc_msg"`
}
