package item

import (
	"dungeonWhiteBox/src/dice"
	"dungeonWhiteBox/src/money"
)

type Item struct {
	Name        string
	Weight      uint16
	Value       money.Coins
	IsBroken    bool
	Tags        []string
	Description string
}

type Weapon struct {
	Base       Item
	Damage     dice.DicePool
	Range      uint8
	IsTwoHand  bool
	IsUsePrist bool
}

type Armor struct {
	Base Item
	Ac   uint8
}
