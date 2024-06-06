package models

import "errors"

// Err - структур для представления ошибок
type Err struct {
	Source  string
	Message string `json:"message"`
}

// Ошибки при работе с ссылками в БД
var (
	ErrAlreadyExists  = errors.New("URL already exists")
	ErrGetDeletedLink = errors.New("deleted Link cant be retrieved")
	ErrGenerateCookie = errors.New("cant generate cookie")
)
