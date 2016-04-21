// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package table

import (
	//"fmt"
	"reflect"
	"testing"

	"encoding/json"

	"fmt"

	"github.com/gotstago/todoapp-go-es/common"
)

type tableTest struct {
	name   string
	input  <-chan common.CommandMessage
	output []common.EventMessage
}

// collect gathers the emitted events into a slice.
func collect(t *tableTest) (output []common.EventMessage) {
	l := parseCommands(t.name, t.input)
	for {
		event := l.nextEvent()
		output = append(output, event)
		if event.Typ == common.MessageEOG || event.Typ == common.MessageError {
			break
		}
	}
	return
}

func equal(i1, i2 []common.EventMessage) bool {
	if len(i1) != len(i2) {
		fmt.Println("i1 type is  :: ", i1[0].Typ, i1[1].Typ)
		fmt.Println("len in equal :: ", len(i1), len(i2))
		return false
	}

	for k := range i1 {
		if i1[k].Typ != i2[k].Typ {
			fmt.Println("types in equal :: ", i1[k].Typ, i2[k].Typ)
			return false
		}
		// var target json.RawMessage
		// source := []byte(*i1[k].Data)
		// err := json.Unmarshal(source, &target)
		// if err != nil {
		// 	fmt.Println("Error %s", err.Error())
		// }
		//fmt.Println("result of unmarshalled data...", string(target))
		if i1[k].Data != nil {
			var e1 interface{}
			if err := json.Unmarshal([]byte(*i1[k].Data), &e1); err != nil {
				fmt.Println(err)
			}
			fmt.Printf("e1 is %v", e1)
		}
		if !reflect.DeepEqual(i1[k].Data, i2[k].Data) {
			fmt.Println("data in equal :: ", i1[k].Data, i2[k].Data)
			return false
		}
		if i1[k].Name != i2[k].Name {
			fmt.Println("name in equal :: ", i1[k].Name, i2[k].Name)
			return false
		}

		if fromJSON(i1[k].Data) != fromJSON(i2[k].Data) {
			//fmt.Println("name in equal :: ", i1[k].Name, i2[k].Name)
			return false
		}
	}
	return true
}

func fromJSON(d *json.RawMessage) interface{} {
    if d == nil {
        return nil
    }
	var d1 interface{}
	if err := json.Unmarshal([]byte(*d), &d); err != nil {
		fmt.Println(err)
        return err
	}
    fmt.Println("returning ", d1)
    return d1
}

//
func toChannel(actions []common.CommandMessage) <-chan common.CommandMessage {
	out := make(chan common.CommandMessage)
	go func() { //goroutine allows code to block but does not block main thread
		for _, n := range actions {
			// fmt.Println("before writing CommandMessage ")
			out <- n
			// fmt.Println("after writing CommandMessage ")
		}
		close(out)
	}()
	return out
}

//for unmarshalling rawJson,
//see https://www.socketloop.com/tutorials/golang-marshal-and-unmarshal-json-rawmessage-struct-example
var (
	cmdData    = []byte(`{"command":"ping","repeat":true}`)
	cmdError   = common.CommandMessage{Name: "", Data: nil, Typ: common.MessageError}
	eventName  = "ping"
	eventError = common.EventMessage{
		Name: "",
		Typ:  common.MessageError,
		Data: nil,
	}
	data, _ = json.Marshal(struct {
		command string
		repeat  bool
	}{
		"ping",
		true,
	})
)

// Some easy cases

var tableEasyTests = []tableTest{
	{"start",
		toChannel([]common.CommandMessage{
			common.CommandMessage{
				Name: "ping",
				Data: (*json.RawMessage)(&cmdData),
				Typ:  common.MessageBid,
			},
		}),
		[]common.EventMessage{
			common.EventMessage{
				Name: eventName,
				Typ:  common.MessageBid,
				Data: (*json.RawMessage)(&cmdData),
			},
			eventError,
		},
	},
	// {"empty action", `$$@@`, []item{tLeftDelim, tRightDelim, tEOF}},
	// {"for", `$$for@@`, []item{tLeftDelim, tFor, tRightDelim, tEOF}},
	// {"quote", `$$"abc \n\t\" "@@`, []item{tLeftDelim, tQuote, tRightDelim, tEOF}},
	// {"raw quote", "$$" + raw + "@@", []item{tLeftDelim, tRawQuote, tRightDelim, tEOF}},
}

func TestTable(t *testing.T) {
	//t.Parallel()
	for _, test := range tableEasyTests {
		output := collect(&test)
		if !equal(output, test.output) {
			t.Errorf("%s: got\n\t%+v\nexpected\n\t%+v", test.name, output, test.output)
			//t.Error("error")
		}
	}
}
