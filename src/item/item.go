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
	Tags        map[string]int32
	Description string
}

type BoostItem struct {
	Base          Item
	PropertyConst map[string]uint32
	PropertyDice  map[string]dice.DicePool
	IsNegative    bool
	ActionTime    uint8
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
