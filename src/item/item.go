package item

import "dungeonWhiteBox/src/money"

type Item struct {
	Name        string
	Weight      uint16
	Value       money.Coins
	IsBroken    bool
	Tags        []string
	Description string
}
