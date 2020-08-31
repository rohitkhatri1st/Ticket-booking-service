package models

type GenericResponse struct {
	Message    string
	Data       interface{}
	StatusCode int
}
