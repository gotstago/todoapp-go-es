package game

import (
	"encoding/json"
	"log"

	"github.com/gotstago/todoapp-go-es/common"
	"github.com/gotstago/todoapp-go-es/event"
	"github.com/gotstago/todoapp-go-es/fsstore"
)

//Projection the game projection which creates game views
type Projection struct {
	subscription *event.Subscription
	datastore    fsstore.FSStore
}

//NewProjection creates a new Projection
func NewProjection(bus event.Bus) *Projection {
	datastore, err := fsstore.NewJSONFSStore("game")
	if err != nil {
		panic(err)
	}
	p := &Projection{
		subscription: bus.Subscribe(
			"GameProjection",
			eventGameItemCreated,
			eventGameItemRemoved,
			eventGameItemUpdated,
		),
		datastore: datastore,
	}

	go p.start()

	return p
}

//HandleEvent handles events this projection subscribes to
func (p *Projection) HandleEvent(event *common.EventMessage) {
	switch event.Name {
	case eventGameItemUpdated:
		fallthrough//skips the next case evaluation and executes 45
	case eventGameItemCreated:
		p.handleGameItemCreatedEvent(event)
	case eventGameItemRemoved:
		p.handleGameItemRemovedEvent(event)
	}
}

func (p *Projection) handleGameItemCreatedEvent(event *common.EventMessage) {
    log.Println("about to handleGameItemCreatedEvent.....")
	game := new(Game)
	err := json.Unmarshal(*event.Data, game)
	if err != nil {
		panic(err)
	}
	p.datastore.Set(game.ID, game)
	p.datastore.AddToCollection("all", game.ID, game)
}

func (p *Projection) handleGameItemRemovedEvent(event *common.EventMessage) {
	var id string
	err := json.Unmarshal(*event.Data, &id)
	if err != nil {
		log.Panic(err)
	}
	p.datastore.Remove(id)
	p.datastore.RemoveFromCollection("all", id)
}

func (p *Projection) start() {
	for {
		select {
		case event := <-p.subscription.EventChan:
			//go - am - think this was a bug, results in events getting handled out of turn
            log.Println("about to handle event.....")
            p.HandleEvent(event)
		}
	}
}
