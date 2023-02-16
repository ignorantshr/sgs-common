package enum

type CardType int

const (
	CardTypeBase CardType = iota + 1
)

type CardSuitType string

const (
	CardSuitTypeHeart   = "Heart"   // 红桃
	CardSuitTypeDiamond = "Diamond" // 方块
	CardSuitTypeClub    = "Club"    // 梅花
	CardSuitTypeSpade   = "Spade"   // 黑桃
)

type CardColorType string

const (
	CardColorTypeRed   = "Red"
	CardColorTypeBlack = "Black"
)

func GetColorType(suit CardSuitType) CardColorType {
	var color CardColorType
	switch suit {
	case CardSuitTypeClub, CardSuitTypeSpade:
		color = CardColorTypeBlack
	case CardSuitTypeDiamond, CardSuitTypeHeart:
		color = CardColorTypeRed
	}
	return color
}

type CardName string

const (
	CardNameKill  = "杀"
	CardNameDodge = "闪"
)
