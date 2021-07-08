package beer

type Beer struct {
	ID    int64     `json:"id"`
	Name  string    `json:"name"`
	Type  BeerType  `json:"type"`
	Style BeerStyle `json:"style"`
}

type BeerType int

const (
	TypeAle = iota + 1
	TypeLager
	TypeMalt
	TypeStout
)

func (bt BeerType) String() string {
	switch bt {
	case TypeAle:
		return "Ale"
	case TypeLager:
		return "Lager"
	case TypeMalt:
		return "Malt"
	case TypeStout:
		return "Stout"
	}
	return "Unknow"
}

type BeerStyle int

const (
	StyleAmber = iota + 1
	StyleBlonde
	StyleBrown
	StyleCream
	StyleDark
	StylePale
	StyleStrong
	StyleWheat
	StyleRed
	StyleIPA
	StyleLime
	StylePilsner
	StyleGolden
	StyleFruit
	StyleHoney
)

func (bs BeerStyle) String() string {
	switch bs {
	case StyleAmber:
		return "Amber"
	case StyleBlonde:
		return "Blonde"
	case StyleBrown:
		return "Brown"
	case StyleCream:
		return "Cream"
	case StyleDark:
		return "Dark"
	case StylePale:
		return "Pale"
	case StyleStrong:
		return "Strong"
	case StyleWheat:
		return "Wheat"
	case StyleRed:
		return "Red"
	case StyleIPA:
		return "IPA"
	case StyleLime:
		return "Lime"
	case StylePilsner:
		return "Pilsner"
	case StyleGolden:
		return "Golden"
	case StyleFruit:
		return "Fruit"
	case StyleHoney:
		return "Honey"
	}
	return "Unknow"
}
