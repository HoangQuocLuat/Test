package models

type Response struct {
	StatusCode int
	Message string
	Data interface{}
}