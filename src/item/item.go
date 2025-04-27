package item

import "dungeonWhiteBox/src/money"

type Item struct {
	name        string
	weight      float32
	coins       *money.Coins
	tags        []string
	Description string
}
