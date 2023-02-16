package sgs_common

import "github.com/ignorantshr/sgs-common/enum"

type Event struct {
	Type        enum.EventType           `json:"type,omitempty"`
	InOut       enum.EventInOutType      `json:"in_out,omitempty"`
	Visibility  enum.EventVisibilityType `json:"visibility,omitempty"`
	GameId      string                   `json:"game_id,omitempty"`
	InitiatorId string                   `json:"initiator_id,omitempty"`
	ReceiversId []string                 `json:"receivers_id,omitempty"`
	ExtraInfo   map[string]interface{}   `json:"extra_info,omitempty"`
}

func NewEmptyEvent() *Event {
	return &Event{}
}

func NewEvent(types enum.EventType, inOut enum.EventInOutType, visibility enum.EventVisibilityType, gameId string, initiatorId string, extraInfo map[string]interface{}) *Event {
	return &Event{
		Type:        types,
		InOut:       inOut,
		Visibility:  visibility,
		GameId:      gameId,
		InitiatorId: initiatorId,
		ReceiversId: make([]string, 0),
		ExtraInfo:   extraInfo,
	}
}

func NewNoticeEvent(gameId string, initiatorId string, extraInfo map[string]interface{}) *Event {
	return &Event{
		Type:        enum.EventTypeNotice,
		InOut:       enum.EventInOutTypeOut,
		Visibility:  enum.EventVisibilityTypeAll,
		GameId:      gameId,
		InitiatorId: initiatorId,
		ReceiversId: make([]string, 0),
		ExtraInfo:   extraInfo,
	}
}

func (e *Event) AddReceiverList(list []string) {
	e.ReceiversId = append(e.ReceiversId, list...)
}

func (e *Event) IsReceiver(playerId string) bool {
	for _, id := range e.ReceiversId {
		if id == playerId {
			return true
		}
	}
	return false
}

func (e *Event) IsIn() bool {
	return e.InOut == enum.EventInOutTypeIn
}

func (e *Event) IsOut() bool {
	return e.InOut == enum.EventInOutTypeOut
}
