package main

import (
	"dungeonWhiteBox/src/dice"
	"fmt"
)

func main() {
	roll := dice.Roll(20)
	fmt.Println("d20:", roll)
	roll = dice.Roll(20)
	fmt.Println("d20:", roll)
	roll = dice.Roll(20)
	fmt.Println("d20:", roll)
}
