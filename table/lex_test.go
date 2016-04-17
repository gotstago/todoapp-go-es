// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package table

import (
	"fmt"
	"testing"

	"encoding/json"

	"github.com/gotstago/todoapp-go-es/common"
)

// Make the types prettyprint.
// var itemName = map[itemType]string{
// 	itemError:        "error",
// 	itemBool:         "bool",
// 	itemChar:         "char",
// 	itemCharConstant: "charconst",
// 	itemComplex:      "complex",
// 	itemColonEquals:  ":=",
// 	itemEOF:          "EOF",
// 	itemField:        "field",
// 	itemIdentifier:   "identifier",
// 	itemLeftDelim:    "left delim",
// 	itemLeftParen:    "(",
// 	itemNumber:       "number",
// 	itemPipe:         "pipe",
// 	itemRawString:    "raw string",
// 	itemRightDelim:   "right delim",
// 	itemRightParen:   ")",
// 	itemSpace:        "space",
// 	itemString:       "string",
// 	itemVariable:     "variable",

// 	// keywords
// 	itemDot:      ".",
// 	itemBlock:    "block",
// 	itemDefine:   "define",
// 	itemElse:     "else",
// 	itemIf:       "if",
// 	itemEnd:      "end",
// 	itemNil:      "nil",
// 	itemRange:    "range",
// 	itemTemplate: "template",
// 	itemWith:     "with",
// }

// func (i itemType) String() string {
// 	s := itemName[i]
// 	if s == "" {
// 		return fmt.Sprintf("item%d", int(i))
// 	}
// 	return s
// }

type tableTest struct {
	name   string
	input  <-chan common.CommandMessage
	output []common.EventMessage
}

// var (
// 	tDot        = item{itemDot, 0, "."}
// 	tBlock      = item{itemBlock, 0, "block"}
// 	tEOF        = item{itemEOF, 0, ""}
// 	tFor        = item{itemIdentifier, 0, "for"}
// 	tLeft       = item{itemLeftDelim, 0, "{{"}
// 	tLpar       = item{itemLeftParen, 0, "("}
// 	tPipe       = item{itemPipe, 0, "|"}
// 	tQuote      = item{itemString, 0, `"abc \n\t\" "`}
// 	tRange      = item{itemRange, 0, "range"}
// 	tRight      = item{itemRightDelim, 0, "}}"}
// 	tRpar       = item{itemRightParen, 0, ")"}
// 	tSpace      = item{itemSpace, 0, " "}
// 	raw         = "`" + `abc\n\t\" ` + "`"
// 	rawNL       = "`now is{{\n}}the time`" // Contains newline inside raw quote.
// 	tRawQuote   = item{itemRawString, 0, raw}
// 	tRawQuoteNL = item{itemRawString, 0, rawNL}
// )
var (
	cmdName        = "ping"
	cmdData        = []byte(`{"command": "ping"}`)
	cmdRawData     = (*json.RawMessage)(&cmdData)
	cmdType        = common.CommandAnnounce
	cmdSource      = common.CommandMessage{Name: cmdName, Data: cmdRawData, Typ: cmdType}
	cmdSourceSlice = []common.CommandMessage{cmdSource}
)

// Some easy cases
var tableEasyTests = []tableTest{
	{"start", gen(cmdSourceSlice), []common.EventMessage{}},
	// {"empty action", `$$@@`, []item{tLeftDelim, tRightDelim, tEOF}},
	// {"for", `$$for@@`, []item{tLeftDelim, tFor, tRightDelim, tEOF}},
	// {"quote", `$$"abc \n\t\" "@@`, []item{tLeftDelim, tQuote, tRightDelim, tEOF}},
	// {"raw quote", "$$" + raw + "@@", []item{tLeftDelim, tRawQuote, tRightDelim, tEOF}},
}

func gen(actions []common.CommandMessage) <-chan common.CommandMessage {
	out := make(chan common.CommandMessage)
	go func() { //goroutine allows code to block but does not block main thread
		for _, n := range actions {
			out <- n
			fmt.Println("writing CommandMessage ", n)
		}
		close(out)
	}()
	return out
}

// collect gathers the emitted events into a slice.
func collect(t *tableTest) (output []common.EventMessage) {
	l := parseCommands(t.name, t.input)
	for {
		event := l.nextItem()
		output = append(output, event)
		if event.Typ == common.EventEOG || event.Typ == common.EventError {
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
		if i1[k].Typ != i2[k].Typ {
			return false
		}
		if i1[k].Data != i2[k].Data {
			return false
		}
	}
	return true
}

func TestTable(t *testing.T) {
	for _, test := range tableEasyTests {
		output := collect(&test)
		if !equal(output, test.output) {
			t.Errorf("%s: got\n\t%+v\nexpected\n\t%v", test.name, output, test.output)
		}
	}
}

// Some easy cases from above, but with delimiters $$ and @@
// var lexDelimTests = []tableTest{
// 	{"punctuation", "$$,@%{{}}@@", []item{
// 		tLeftDelim,
// 		{itemChar, 0, ","},
// 		{itemChar, 0, "@"},
// 		{itemChar, 0, "%"},
// 		{itemChar, 0, "{"},
// 		{itemChar, 0, "{"},
// 		{itemChar, 0, "}"},
// 		{itemChar, 0, "}"},
// 		tRightDelim,
// 		tEOF,
// 	}},
// 	{"empty action", `$$@@`, []item{tLeftDelim, tRightDelim, tEOF}},
// 	{"for", `$$for@@`, []item{tLeftDelim, tFor, tRightDelim, tEOF}},
// 	{"quote", `$$"abc \n\t\" "@@`, []item{tLeftDelim, tQuote, tRightDelim, tEOF}},
// 	{"raw quote", "$$" + raw + "@@", []item{tLeftDelim, tRawQuote, tRightDelim, tEOF}},
// }

// var (
// 	tLeftDelim  = item{itemLeftDelim, 0, "$$"}
// 	tRightDelim = item{itemRightDelim, 0, "@@"}
// )

// func TestDelims(t *testing.T) {
// 	for _, test := range lexDelimTests {
// 		output := collect(&test, "$$", "@@")
// 		if !equal(output, test.output, false) {
// 			t.Errorf("%s: got\n\t%v\nexpected\n\t%v", test.name, output, test.output)
// 		}
// 	}
// }

// var lexPosTests = []tableTest{
// 	{"empty", "", []item{tEOF}},
// 	{"punctuation", "{{,@%#}}", []item{
// 		{itemLeftDelim, 0, "{{"},
// 		{itemChar, 2, ","},
// 		{itemChar, 3, "@"},
// 		{itemChar, 4, "%"},
// 		{itemChar, 5, "#"},
// 		{itemRightDelim, 6, "}}"},
// 		{itemEOF, 8, ""},
// 	}},
// 	{"sample", "0123{{hello}}xyz", []item{
// 		{itemText, 0, "0123"},
// 		{itemLeftDelim, 4, "{{"},
// 		{itemIdentifier, 6, "hello"},
// 		{itemRightDelim, 11, "}}"},
// 		{itemText, 13, "xyz"},
// 		{itemEOF, 16, ""},
// 	}},
// }

// // The other tests don't check position, to make the test cases easier to construct.
// // This one does.
// func TestPos(t *testing.T) {
// 	for _, test := range lexPosTests {
// 		output := collect(&test, "", "")
// 		if !equal(output, test.output, true) {
// 			t.Errorf("%s: got\n\t%v\nexpected\n\t%v", test.name, output, test.output)
// 			if len(output) == len(test.output) {
// 				// Detailed print; avoid item.String() to expose the position value.
// 				for i := range output {
// 					if !equal(output[i:i+1], test.output[i:i+1], true) {
// 						i1 := output[i]
// 						i2 := test.output[i]
// 						t.Errorf("\t#%d: got {%v %d %q} expected  {%v %d %q}", i, i1.typ, i1.pos, i1.val, i2.typ, i2.pos, i2.val)
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// Test that an error shuts down the lexing goroutine.
// func TestShutdown(t *testing.T) {
// 	// We need to duplicate template.Parse here to hold on to the lexer.
// 	const text = "erroneous{{define}}{{else}}1234"
// 	lexer := lex("foo", text, "{{", "}}")
// 	_, err := New("root").parseLexer(lexer, text)
// 	if err == nil {
// 		t.Fatalf("expected error")
// 	}
// 	// The error should have drained the input. Therefore, the lexer should be shut down.
// 	token, ok := <-lexer.output
// 	if ok {
// 		t.Fatalf("input was not drained; got %v", token)
// 	}
// }

// // parseLexer is a local version of parse that lets us pass in the lexer instead of building it.
// // We expect an error, so the tree set and funcs list are explicitly nil.
// func (t *Tree) parseLexer(lex *lexer, text string) (tree *Tree, err error) {
// 	defer t.recover(&err)
// 	t.ParseName = t.Name
// 	t.startParse(nil, lex, map[string]*Tree{})
// 	t.parse()
// 	t.add()
// 	t.stopParse()
// 	return t, nil
// }
