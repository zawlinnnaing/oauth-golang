package app_error

import json "encoding/json"

type ValidationField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type ValidationError struct {
	Errors []ValidationField `json:"errors"`
}

func (e *ValidationError) Set(field string, message string) {
	e.Errors = append(e.Errors, ValidationField{Field: field, Message: message})
}

func (e *ValidationError) Error() string {
	s, _ := json.Marshal(*e)
	return string(s)
}

func NewValidationError() *ValidationError {
	return &ValidationError{}
}
