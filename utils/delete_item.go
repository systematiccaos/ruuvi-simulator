package utils

// Deletes item at index in a slice
func DeleteItem[T any](slice []T, index int) []T {
	return append(slice[:index], slice[index+1:]...)
}
