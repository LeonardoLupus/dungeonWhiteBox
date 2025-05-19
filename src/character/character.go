package character

import (
	"dungeonWhiteBox/src/item"
)

var Attributes = map[string]uint32{
	"Strength":     0,
	"Intelligence": 0,
	"Wisdom":       0,
	"Constitution": 0,
	"Dexterity":    0,
	"Charisma":     0,
}

type character struct {
	Name       string
	attributes map[string]uint32
	class      ClassId
	race       RaceId
	ac         uint32
	hp         uint32
	backpack   []item.InventoryItem
}
