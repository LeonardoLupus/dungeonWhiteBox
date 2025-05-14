package character

type RaceId int

const (
	emptyRace RaceId = iota
	human
	dwarf
	elf
	halfling
)

var Races = map[RaceId]string{
	emptyRace: "EMPTY",
	human:     "Human",
	dwarf:     "Dwarf",
	elf:       "Elf",
	halfling:  "Halfling",
}
