package utils

import "math/rand"

func GenerateRandomIntInRange(
	min,
	max int,
) int {
	return rand.Intn(max-min+1) + min
}
