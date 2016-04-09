package game

import (
	"encoding/json"

	"github.com/pborman/uuid"

	"github.com/gotstago/todoapp-go-es/common"
)

const (
	eventGameItemCreated = "gameItemCreated"
	eventGameItemRemoved = "gameItemRemoved"
	eventGameItemUpdated = "gameItemUpdated"
)

//CreateGameItem creates a game based on a command message
func CreateGameItem(cmd *common.CommandMessage, eventChan chan<- *common.EventMessage) error {
	var game Game
	if err := json.Unmarshal(*cmd.Data, &game); err != nil {
		return err
	}
	game.ID = uuid.New()

	data, err := json.Marshal(game)
	if err != nil {
		return err
	}

	raw := json.RawMessage(data)

	event := &common.EventMessage{
		Name: eventGameItemCreated,
		Data: &raw,
	}
	eventChan <- event
	return nil
}

//RemoveGameItem removes a game based on a command message
func RemoveGameItem(cmd *common.CommandMessage, eventChan chan<- *common.EventMessage) error {
	event := &common.EventMessage{
		Name: eventGameItemRemoved,
		Data: cmd.Data,
	}
	eventChan <- event
	return nil
}

//UpdateGameItem updates a game based on a command message
func UpdateGameItem(cmd *common.CommandMessage, eventChan chan<- *common.EventMessage) error {
	event := &common.EventMessage{
		Name: eventGameItemUpdated,
		Data: cmd.Data,
	}
	eventChan <- event
	return nil
}
