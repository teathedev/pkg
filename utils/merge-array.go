package utils

func MergeArrays[T any](arrays ...[]T) []T {
	var merged []T

	for _, arr := range arrays {
		merged = append(merged, arr...)
	}

	return merged
}
