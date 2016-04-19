package common

import (
	"encoding/json"
	"fmt"
)

//CommandMessage is a WS command message
//see https://willnorris.com/2014/05/go-rest-apis-and-pointers for info on pointers
type CommandMessage struct {
	Name string           `json:"name"     validate:"nonzero"`
	Data *json.RawMessage `json:"data"`//defaults to nil
	Typ  MessageType      `json:"type"`
}

// MessageType identifies the type of incoming command.
type MessageType int

const (
	MessageError    MessageType = iota // error occurred; value is text of error
	MessageBid                         // player bid
	MessageAnnounce                    // player announcement - eg. Bella
	MessagePlayCard                    // player submitting a card to play
	MessageDeal                        // player request to deal
	MessageAccuse                      // accuse another player of a misplay
	MessageEOG                         //end of game
)

func (cm CommandMessage) String() string {
	var data string
    if err := json.Unmarshal(*cm.Data, &data); err != nil {
		return err.Error()
	}
    switch {
	case cm.Typ == MessageEOG:
		return "EOG"
	case cm.Typ == MessageError:
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
	Typ     MessageType        `json:"type"`
}

// EventType identifies the type of Event.
// type EventType int

// const (
// 	EventError    EventType = iota // EventError occurred; value is text of error
// 	EventBid                       // player bid
// 	EventAnnounce                  // player announcement - eg. Bella
// 	EventPlayCard                  // player submitting a card to play
// 	EventDeal                      // player request to deal
// 	EventAccuse                    // accuse another player of a misplay
// 	EventEOG                       //end of game
// )

func (em EventMessage) String() string {
	var data json.RawMessage
    if err := json.Unmarshal(*em.Data, &data); err != nil {
		return err.Error()
	}
    switch {
	case em.Typ == MessageEOG:
		return "EOG"
	case em.Typ == MessageError:
		return string(data)
	case len(data) > 10:
		return fmt.Sprintf("%.10q...", data)
	}
	return fmt.Sprintf("%q", data)
}
