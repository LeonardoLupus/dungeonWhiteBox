package dice

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

var ErrInvalidDises = errors.New("invalid number of dice sides")

func errorMessage(message string) {
	fmt.Printf("\n[Error] %s\n", message)
}

func isCorrectCountOrSidesDices(count, sides int32) bool {
	if sides <= 0 || count <= 0 {
		errorMessage("Dises without sides or count dises is below zero!")
		return false
	}
	return true
}

func Roll(sides int32) (int32, error) {
	if !isCorrectCountOrSidesDices(1, sides) {
		return -1, ErrInvalidDises
	}
	return rand.Int32N(sides) + 1, nil
}

func RollMultiple(count, sides int32) (int32, error) {
	if !isCorrectCountOrSidesDices(count, sides) {
		return -1, ErrInvalidDises
	}

	var sum int32
	for range count {
		sum += rand.Int32N(sides) + 1
	}
	return sum, nil
}

func RollBool() bool {
	return rand.IntN(2) == 0
}

type DicePool struct {
	Count    int32
	Sides    int32
	Modifier int32
}

func (d *DicePool) SetDicePool(count, sides, modifier int32) {
	if !isCorrectCountOrSidesDices(count, sides) {
		return
	}
	d.Count = count
	d.Sides = sides
	d.Modifier = modifier
}

func (d DicePool) GetDicePool() (int32, int32, int32) {
	return d.Count, d.Sides, d.Modifier
}

func (d DicePool) Roll() (int32, error) {
	var res, err = RollMultiple(d.Count, d.Sides)
	if err != nil {
		return -1, ErrInvalidDises
	}
	res += d.Modifier
	if res < 0 {
		res = 0
	}
	return res, nil
}
