package beer

type Beer struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Type  Type   `json:"type"`
	Style Style  `json:"style"`
}

type Type int

const (
	Ale   = 1
	Lager = 2
	Malt  = 3
	Stout = 4
)

func (beerType Type) AsString() string {
	switch beerType {
	case Ale:
		return "Ale"
	case Lager:
		return "Lager"
	case Malt:
		return "Malt"
	case Stout:
		return "Stout"
	}

	return "Unknown"
}

type Style string

const (
	Amber = iota + 1
	Blonde
	Brown
	Cream
	Dark
	Pale
	Strong
	Wheat
	Red
	IPA
	Lime
	Pilsner
	Golden
	Fruit
	Honey
)

func (style Style) AsString() string {
	switch style {
	case Amber:
		return "Amber"
	case Blonde:
		return "Blonde"
	case Brown:
		return "Brown"
	case Cream:
		return "Cream"
	case Dark:
		return "Dark"
	case Pale:
		return "Pale"
	case Strong:
		return "Strong"
	case Wheat:
		return "Wheat"
	case Red:
		return "Red"
	case IPA:
		return "India Pale Ale"
	case Lime:
		return "Lime"
	case Pilsner:
		return "Pilsner"
	case Golden:
		return "Golden"
	case Fruit:
		return "Fruit"
	case Honey:
		return "Honey"
	}

	return "Unknown"
}
