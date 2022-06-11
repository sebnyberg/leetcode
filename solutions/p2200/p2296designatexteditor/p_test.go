package p2296designatexteditor

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTextEditor(t *testing.T) {
	type testAction struct {
		actName string
		args    []any
		want    []any
	}

	const (
		actNameAddText     = "addText"
		actNameCursorLeft  = "cursorLeft"
		actNameDeleteText  = "deleteText"
		actNameCursorRight = "cursorRight"
	)

	for i, tc := range [][]testAction{
		{
			{actNameAddText, []any{"bxyackuncqzcqo"}, nil},
			{actNameCursorLeft, []any{12}, []any{"bx"}},
			{actNameDeleteText, []any{3}, []any{2}},
			{actNameCursorLeft, []any{5}, []any{""}},
			{actNameAddText, []any{"osdhyvqxf"}, nil},
			{actNameCursorRight, []any{10}, []any{"yackuncqzc"}},
		},
	} {
		t.Run(fmt.Sprintf("test: %d", i), func(t *testing.T) {
			a := Constructor()
			for j, act := range tc {
				switch act.actName {
				case actNameAddText:
					a.AddText(act.args[0].(string))
				case actNameDeleteText:
					res := a.DeleteText(act.args[0].(int))
					require.Equal(t, act.want[0].(int), res, fmt.Sprintf("action %d, name: %s", j, act.actName))
				case actNameCursorLeft:
					res := a.CursorLeft(act.args[0].(int))
					require.Equal(t, act.want[0].(string), res, fmt.Sprintf("action %d, name: %s", j, act.actName))
				case actNameCursorRight:
					res := a.CursorRight(act.args[0].(int))
					require.Equal(t, act.want[0].(string), res, fmt.Sprintf("action %d, name: %s", j, act.actName))

				default:
					panic(act.actName)
				}
			}
		})
	}
}

type TextEditor struct {
	text      []byte
	cursorPos int
}

func Constructor() TextEditor {
	return TextEditor{
		text:      make([]byte, 0, 100),
		cursorPos: 0,
	}
}

func (this *TextEditor) AddText(text string) {
	k := len(text)
	end := len(this.text) + k
	if cap(this.text) >= end {
		this.text = this.text[:len(this.text)+k]
	} else {
		d := end - cap(this.text)
		this.text = this.text[:cap(this.text)]
		this.text = append(this.text, make([]byte, d)...)
	}
	copy(this.text[this.cursorPos+k:], this.text[this.cursorPos:])
	copy(this.text[this.cursorPos:], text)
	this.cursorPos += k
}

func (this *TextEditor) DeleteText(k int) int {
	a := k - this.cursorPos
	if a > 0 {
		k -= a
	}
	newCursorPos := this.cursorPos - k
	copy(this.text[newCursorPos:], this.text[this.cursorPos:])
	this.text = this.text[:len(this.text)-k]
	this.cursorPos = newCursorPos
	return k
}

func (this *TextEditor) CursorLeft(k int) string {
	this.cursorPos = max(0, this.cursorPos-k)
	start := max(0, this.cursorPos-10)
	s := string(this.text[start:this.cursorPos])
	return s
}

func (this *TextEditor) CursorRight(k int) string {
	this.cursorPos = min(len(this.text), this.cursorPos+k)
	start := max(0, this.cursorPos-10)
	s := string(this.text[start:this.cursorPos])
	return s
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
