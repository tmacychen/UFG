// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package framework

type TextBox interface {
	//j	Focusable
	//j	OnSelectionChanged(func()) EventSubscription
	//j	OnTextChanged(func([]TextBoxEdit)) EventSubscription
	//j	Padding() math.Spacing
	//j	SetPadding(math.Spacing)
	//j	Runes() []rune
	//j	Text() string
	//j	SetText(string)
	//j	Font() Font
	//j	SetFont(Font)
	//j	Multiline() bool
	//j	SetMultiline(bool)
	//j	DesiredWidth() int
	//j	SetDesiredWidth(desiredWidth int)
	//j	TextColor() Color
	//j	SetTextColor(Color)
	//j	Select(TextSelectionList)
	//j	SelectAll()
	//j	Carets() []int
	//j	RuneIndexAt(p math.Point) (idx int, found bool)
	//j	TextAt(s, e int) string
	//j	WordAt(runeIndex int) string
	//j	ScrollToLine(int)
	//j	ScrollToRune(int)
	//j	LineIndex(runeIndex int) int
	//j	LineStart(line int) int
	//j	LineEnd(line int) int
}
