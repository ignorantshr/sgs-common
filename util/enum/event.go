package enum

type EventInfoKey string

const (
	EventInfoKeyKillSeq = "killSeq"
)

type EventType int

const (
	EventTypeNotice = iota
)

type EventVisibilityType int

const (
	EventVisibilityTypeAll  = iota
	EventVisibilityTypeSelf = iota
)
