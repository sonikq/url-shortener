package models

// Err - структур для представления ошибок
type Err struct {
	Source  string
	Message string `json:"message"`
}
