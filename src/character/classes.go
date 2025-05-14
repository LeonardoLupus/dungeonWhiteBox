package character

type ClassId int

const (
	emptyClass ClassId = iota
	cleric
	fighter
	wizard
	thief
)

var Classes = map[ClassId]string{
	emptyClass: "EMPTY",
	cleric:     "Cleric",
	fighter:    "Fighter",
	wizard:     "Wizard",
	thief:      "Thief",
}
