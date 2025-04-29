package money

type Coins struct {
	count uint64
}

const GoldToSilver = 10
const SilverToCopper = 10

func (c *Coins) Get() (uint64, uint64, uint64) {
	gold := c.count / (GoldToSilver * SilverToCopper)
	silver := (c.count - gold*GoldToSilver) / SilverToCopper
	copper := (c.count - gold*GoldToSilver*SilverToCopper - silver*SilverToCopper)
	return gold, silver, copper
}

func (c *Coins) Set(gold, silver, copper uint64) {
	c.count = gold*GoldToSilver*SilverToCopper + silver*SilverToCopper + copper
}

func (c *Coins) Add(gold, silver, copper uint64) {
	c.count += gold*GoldToSilver*SilverToCopper + silver*SilverToCopper + copper
}

func (c *Coins) Remove(gold, silver, copper uint64) uint64 {
	sum := (gold*GoldToSilver*SilverToCopper + silver*SilverToCopper + copper)
	if sum > c.count {
		return sum - c.count
	}
	c.count -= sum
	return 0
}

func (c *Coins) Pay(thing *Coins) uint64 {
	if thing.count > c.count {
		return thing.count - c.count
	}
	c.count -= thing.count
	return 0
}

func (c *Coins) Gain(thing *Coins) {
	c.count += thing.count
}
