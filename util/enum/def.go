package enum

// GCountry 所属势力
type GCountry int

const (
	Country1 GCountry = iota
	Country2
	Country3
)

type ActionType int

const (
	ActionTypeReleaseNotice ActionType = iota
)

const DefaultDrawnCardsNum = 2

type GameType int

const (
	GameType5Man GameType = iota
)
