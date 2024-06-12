package utils

// Ptr -
func Ptr[T any](value T) *T {
	return &value
}
