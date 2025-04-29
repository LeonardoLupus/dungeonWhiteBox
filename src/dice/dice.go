package dice

import (
	"fmt"
	"math/rand/v2"
)

func errorMessage(message string) {
	fmt.Printf("\n[Error] %s\n", message)
}

func Roll(sides uint32) uint32 {
	if sides == 0 {
		errorMessage("Dise without sides!")
		return 0
	}
	return rand.Uint32N(sides) + 1
}

func RollMultiple(count, sides uint32) uint32 {
	if sides == 0 || count == 0 {
		errorMessage("Dises without sides or count dises is zero!")
		return 0
	}

	var sum uint32
	for range count {
		sum += rand.Uint32N(sides) + 1
	}
	return sum
}

func RollBool() bool {
	return rand.IntN(2) == 0
}

type DicePool struct {
	count    uint32
	sides    uint32
	modifier int32
}

func (d *DicePool) SetDicePool(count, sides uint32, modifier int32) {
	d.count = count
	d.sides = sides
	d.modifier = modifier
}

func (d *DicePool) GetDicePool() (uint32, uint32, int32) {
	return d.count, d.sides, d.modifier
}

func (d *DicePool) Roll() uint32 {
	if d.count == 0 || d.sides == 0 {
		errorMessage("Dises without sides or count dises is zero!")
		return 0
	}
	var res = RollMultiple(d.count, d.sides)
	if d.modifier < 0 {
		d.modifier *= -1
		if uint32(d.modifier) >= res {
			res = 0
		} else {
			res -= uint32(d.modifier)
		}
	} else {
		res += uint32(d.modifier)
	}
	return res
}
