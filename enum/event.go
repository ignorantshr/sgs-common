package enum

type EventInfoKey string

const (
	EventInfoKeyKillSeq = "killSeq"
	EventInfoKeyOptions = "options"
	EventInfoKeyCard    = "card"
	EventInfoKeyTargets = "targets"
)

type EventInOutType int

const (
	EventInOutTypeIn = iota
	EventInOutTypeOut
)

type EventType int

const (
	_ EventType = 0

	EventTypeGameStart EventType = 1001
	EventTypeGameEnd   EventType = 1002

	EventTypeStageJudgment EventType = 2001
	EventTypeStageDrawCard EventType = 2002
	EventTypeStagePlay     EventType = 2003
	EventTypeStageEnd      EventType = 2004

	EventTypeNotice EventType = 3000

	// may skill

	EventTypePlayCard EventType = 4000
	EventTypeChoice   EventType = 4001
)

func (e EventType) IsSkillType() bool {
	return int(e)&4000 == 4000
}

type EventVisibilityType int

const (
	EventVisibilityTypeAll = iota
	EventVisibilityTypeSelf
)
