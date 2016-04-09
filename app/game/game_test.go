package game

import (
	"encoding/json"
	"testing"

	"github.com/pborman/uuid"

	"github.com/gotstago/todoapp-go-es/common"
	"github.com/gotstago/todoapp-go-es/event"
)

func TestCreateStack(t *testing.T) {
	//if we do not set the fsstore.DataDir, the default is to place in /tmp on linux, or the os temp dir
    //prevents adding data to real api files
    bus := event.NewDefaultBus()
	projection := NewProjection(bus)
	id := uuid.New()
    g := &Game{ID: id} 
    t.Log("game is ",g.ID)
	data, _ := json.Marshal(g)
    t.Log("game is ",data)
	raw := json.RawMessage(data)
    t.Log("game is ",data)
	e := &common.EventMessage{
		Name:    eventGameItemCreated,
		Data:    &raw,
		Version: 1,
	}
	projection.HandleEvent(e)
}

func TestCreateAndRemoveStack(t *testing.T) {
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
