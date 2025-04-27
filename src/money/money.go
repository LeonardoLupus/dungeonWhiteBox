package money

type Coins struct {
	count uint64
}

const goldToSilver = 10
const silverToCopper = 10

func (c *Coins) Get() (uint64, uint64, uint64) {
	gold := c.count / (goldToSilver * silverToCopper)
	silver := (c.count - gold*goldToSilver) / silverToCopper
	copper := (c.count - gold*goldToSilver*silverToCopper - silver*silverToCopper)
	return gold, silver, copper
}

func (c *Coins) Set(gold, silver, copper uint64) {
	c.count = gold*goldToSilver*silverToCopper + silver*silverToCopper + copper
}

func (c *Coins) Add(gold, silver, copper uint64) {
	c.count += gold*goldToSilver*silverToCopper + silver*silverToCopper + copper
}

func (c *Coins) Remove(gold, silver, copper uint64) uint64 {
	sum := (gold*goldToSilver*silverToCopper + silver*silverToCopper + copper)
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
