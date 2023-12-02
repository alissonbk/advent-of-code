package utils

func contains[T comparable](slice []T, n T) bool {
	for _, v := range slice {
		if v == n {
			return true
		}
	}
	return false
}
