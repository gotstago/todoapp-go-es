package ws

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sync/atomic"

	"github.com/gotstago/todoapp-go-es/common"
	"github.com/gotstago/todoapp-go-es/event"

	"golang.org/x/net/websocket"
)

// Client represents a websocket connection
// the eventBus is subscribed to
// there is a quitChan for stopping activity
type Client struct {
	conn         *websocket.Conn
	eventBus     event.Bus
	subscription *event.Subscription
	quitChan     chan bool
}

const eventNewWebSocketClient = "newWebSocketClient"

var numClient = int32(0)

type subscribe struct {
	Events []string `json:"events"`
}

// NewClient will return a web socket instance pointer
func NewClient(conn *websocket.Conn, eventBus event.Bus) *Client {

	c := &Client{
		conn:         conn,
		eventBus:     eventBus,
		subscription: eventBus.Subscribe(fmt.Sprintf("WS: %s", conn.LocalAddr())),
		quitChan:     make(chan bool, 1),
	}

	return c
}

func (c *Client) sendNumClientsEvent() {
	jsonData, err := json.Marshal(numClient)
	if err != nil {
		panic(err)
	}

	rawData := json.RawMessage(jsonData)

	c.eventBus.Notify(&common.EventMessage{
		Name: eventNewWebSocketClient,
		Data: &rawData,
	})
}

//Listen will increase client count
// and then it will start to listenRead and listenWrite
func (c *Client) Listen() {
	atomic.AddInt32(&numClient, 1)
	c.sendNumClientsEvent()

	go c.listenWrite()
	c.listenRead()
	//listenRead is blocking until Stop is called
	c.subscription.Destroy()
	atomic.AddInt32(&numClient, -1)
	c.sendNumClientsEvent()
}

// listenRead will listen for incoming messages from connection
// incoming messages contain a list of events
// ChangeSubscription is called to indicate an interest in specific events
func (c *Client) listenRead() {
	//s holds a slice of string events
	s := new(subscribe)
	for {
		select {
		case <-c.quitChan:
			c.Stop()
			return
		default:
			err := websocket.JSON.Receive(c.conn, s)
			if err == io.EOF {
				c.Stop()
			} else if err != nil {
				c.Stop()
				log.Fatal(err)
			} else {
				log.Printf("Changing subscription to: %v", s.Events)
				c.subscription.ChangeSubscription(s.Events...)
			}
		}
	}
}

// listenWrite will listen for events of interest on the subscription
// and send them to client connection
func (c *Client) listenWrite() {
	for {
		select {
		case <-c.quitChan:
			c.Stop()
			return
		case message := <-c.subscription.EventChan:
			err := websocket.JSON.Send(c.conn, message)
			if err != nil {
				c.Stop()
				log.Fatal(err)
			}
		}
	}
}

// Stop will send a message to quit channel for all active funcs
func (c *Client) Stop() {
	c.quitChan <- true
}
