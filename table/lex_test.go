// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package table

import (
	//"fmt"
	"reflect"
	"testing"

	"encoding/json"

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
		return false
	}
	for k := range i1 {
		// if i1[k].Typ != i2[k].Typ {
		// 	fmt.Println("types in equal :: ", i1[k].Typ, i2[k].Typ)
		// 	return false
		// }
		// var target json.RawMessage
		// source := []byte(*i1[k].Data)
		// err := json.Unmarshal(source, &target)
		// if err != nil {
		// 	fmt.Println("Error %s", err.Error())
		// }
		//fmt.Println("result of unmarshalled data...", string(target))
		if !reflect.DeepEqual(i1[k], i2[k]) {
			return false
		}
		// if i1[k] != i2[k] {
		// 	return false
		// }
	}
	return true
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

//for unmarshalling rawJson, see https://www.socketloop.com/tutorials/golang-marshal-and-unmarshal-json-rawmessage-struct-example
var (
	cmdName        = "ping"
	cmdData        = []byte(`{"command":"ping","repeat":true}`)
	cmdRawData     = (*json.RawMessage)(&cmdData)
	cmdType        = common.MessageBid
	cmdSource      = common.CommandMessage{Name: cmdName, Data: cmdRawData, Typ: cmdType}
	cmdSourceSlice = []common.CommandMessage{cmdSource}
	eventName      = "ping"
	eventTarget    = common.EventMessage{
		Name: eventName,
		Typ:  common.MessageBid,
		Data: cmdRawData,
	}
)

// Some easy cases
var tableEasyTests = []tableTest{
	{"start", toChannel(cmdSourceSlice), []common.EventMessage{eventTarget}},
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
			t.Errorf("%s: got\n\t%s\nexpected\n\t%s", test.name, output, test.output)
			//t.Error("error")
		}
	}
}
