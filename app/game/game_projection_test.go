package game

import (
	"encoding/json"
	"testing"

	"github.com/pborman/uuid"

	"github.com/gotstago/todoapp-go-es/common"
	"github.com/gotstago/todoapp-go-es/event"
)

func TestCreateGame(t *testing.T) {
	bus := event.NewDefaultBus()
	projection := NewProjection(bus)
	id := uuid.New()
	data, _ := json.Marshal(&Game{ID: id})
	raw := json.RawMessage(data)
	e := &common.EventMessage{
		Name:    eventGameItemCreated,
		Data:    &raw,
		Version: 1,
	}
	projection.HandleEvent(e)
}

func TestCreateAndRemoveGame(t *testing.T) {
	bus := event.NewDefaultBus()
	projection := NewProjection(bus)
	id := uuid.New()
	data, _ := json.Marshal(&Game{ID: id})
	raw := json.RawMessage(data)
	e := &common.EventMessage{
		Name:    eventGameItemCreated,
		Data:    &raw,
		Version: 1,
	}
	projection.HandleEvent(e)

	raw = json.RawMessage(id)
	e = &common.EventMessage{
		Name: eventGameItemRemoved,
		Data: &raw,
	}
}
