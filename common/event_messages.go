package common

import (
	"encoding/json"
	"fmt"
)

//CommandMessage is a WS command message
type CommandMessage struct {
	Name string           `json:"name"`
	Data *json.RawMessage `json:"data"`
	Typ  CommandType      `json:"typ"`
}

// CommandType identifies the type of incoming command.
type CommandType int

const (
	CommandError    CommandType = iota // error occurred; value is text of error
	CommandBid                         // player bid
	CommandAnnounce                    // player announcement - eg. Bella
	CommandPlayCard                    // player submitting a card to play
	CommandDeal                        // player request to deal
	CommandAccuse                      // accuse another player of a misplay
	CommandEOG                         //end of game
)

func (cm CommandMessage) String() string {
	var data string
    if err := json.Unmarshal(*cm.Data, &data); err != nil {
		return err.Error()
	}
    switch {
	case cm.Typ == CommandEOG:
		return "EOG"
	case cm.Typ == CommandError:
		return data
	// case cm.Typ > itemKeyword:
	// 	return fmt.Sprintf("<%s>", cm.val)
	case len(data) > 10:
		return fmt.Sprintf("%.10q...", data)
	}
	return fmt.Sprintf("%q", data)
}

//ErrorMessage is a generic WS error message
type ErrorMessage struct {
	Reason string `json:"reason"`
}

//EventMessage is a WS event message
type EventMessage struct {
	Name    string           `json:"name"`
	Data    *json.RawMessage `json:"data"`
	Version int              `json:"version"`
	Typ     EventType        `json:"typ"`
}

// EventType identifies the type of Event.
type EventType int

const (
	EventError    EventType = iota // EventError occurred; value is text of error
	EventBid                       // player bid
	EventAnnounce                  // player announcement - eg. Bella
	EventPlayCard                  // player submitting a card to play
	EventDeal                      // player request to deal
	EventAccuse                    // accuse another player of a misplay
	EventEOG                       //end of game
)

func (em EventMessage) String() string {
	var data string
    if err := json.Unmarshal(*em.Data, &data); err != nil {
		return err.Error()
	}
    switch {
	case em.Typ == EventEOG:
		return "EOG"
	case em.Typ == EventError:
		return data
	case len(data) > 10:
		return fmt.Sprintf("%.10q...", data)
	}
	return fmt.Sprintf("%q", data)
}
