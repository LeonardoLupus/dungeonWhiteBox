package item

import (
	"dungeonWhiteBox/src/dice"
	"dungeonWhiteBox/src/money"
)

type InventoryItem interface {
	Name() *string
	Weight() *uint32
	Value() *money.Coins
	IsBroken() *bool
	Description() *string
	Tags() *map[string]int32
}

type Item struct {
	name        string
	weight      uint32
	value       money.Coins
	isBroken    bool
	tags        map[string]int32
	description string
}

func (i Item) Name() *string {
	return &i.name
}

func (i Item) Weight() *uint32 {
	return &i.weight
}

func (i Item) Value() *money.Coins {
	return &i.value
}

func (i Item) IsBroken() *bool {
	return &i.isBroken
}

func (i Item) Description() *string {
	return &i.description
}

func (i Item) Tags() *map[string]int32 {
	return &i.tags
}

type BoostItem struct {
	Item
	PropertyConst map[string]uint32
	PropertyDice  map[string]dice.DicePool
	IsNegative    bool
	ActionTime    uint8
}

type Weapon struct {
	Item
	Damage     dice.DicePool
	Range      uint8
	IsTwoHand  bool
	IsUsePrist bool
}

type Armor struct {
	Item
	Ac uint8
}
