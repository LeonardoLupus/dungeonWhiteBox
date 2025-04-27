package dice

import "math/rand/v2"

func Roll(sides int) int {
	if sides <= 0 {
		return -1
	}
	return rand.IntN(sides) + 1
}

func RollMultiple(count, sides int) int {
	if sides <= 0 || count <= 0 {
		return -1
	}

	sum := 0
	for range count {
		sum += Roll(sides)
	}
	return sum
}

func RollBool() bool {
	return rand.IntN(2) == 0
}
