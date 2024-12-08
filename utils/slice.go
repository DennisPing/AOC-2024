package utils

import "errors"

// ReverseIndex retrieves the N-th element from the end of a slice.
// N is 1-based (1 for the last element, 2 for the second to last, etc.).
func ReverseIndex[T any](slice []T, n int) (T, error) {
	var zero T // Zero value for the type T
	if n <= 0 || n > len(slice) {
		return zero, errors.New("index out of bounds")
	}
	return slice[len(slice)-n], nil
}
