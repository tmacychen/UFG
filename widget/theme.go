// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package widget

import (
	"UFG/framework"

	"UFG/widget/tools"
)

type Theme struct {
	DriverInfo               framework.Driver
	DefaultFontInfo          framework.Font
	DefaultMonospaceFontInfo framework.Font

	WindowBackground tools.Color

	BubbleOverlayStyle        Style
	ButtonDefaultStyle        Style
	ButtonOverStyle           Style
	ButtonPressedStyle        Style
	CodeSuggestionListStyle   Style
	DropDownListDefaultStyle  Style
	DropDownListOverStyle     Style
	FocusedStyle              Style
	HighlightStyle            Style
	LabelStyle                Style
	PanelBackgroundStyle      Style
	ScrollBarBarDefaultStyle  Style
	ScrollBarBarOverStyle     Style
	ScrollBarRailDefaultStyle Style
	ScrollBarRailOverStyle    Style
	SplitterBarDefaultStyle   Style
	SplitterBarOverStyle      Style
	TabActiveHighlightStyle   Style
	TabDefaultStyle           Style
	TabOverStyle              Style
	TabPressedStyle           Style
	TextBoxDefaultStyle       Style
	TextBoxOverStyle          Style
}

// framework.Theme compliance
func (t *Theme) Driver() framework.Driver {
	return t.DriverInfo
}

func (t *Theme) DefaultFont() framework.Font {
	return t.DefaultFontInfo
}

func (t *Theme) SetDefaultFont(f framework.Font) {
	t.DefaultFontInfo = f
}

func (t *Theme) DefaultMonospaceFont() framework.Font {
	return t.DefaultMonospaceFontInfo
}

func (t *Theme) SetDefaultMonospaceFont(f framework.Font) {
	t.DefaultMonospaceFontInfo = f
}

//func (t *Theme) CreateBubbleOverlay() framework.BubbleOverlay {
//	return CreateBubbleOverlay(t)
//}
//
//func (t *Theme) CreateButton() framework.Button {
//	return widget.CreateButton(t)
//}
//
//
//func (t *Theme) CreateCodeEditor() framework.CodeEditor {
//	return CreateCodeEditor(t)
//}

//func (t *Theme) CreateDropDownList() framework.DropDownList {
//	return CreateDropDownList(t)
//}
//
//func (t *Theme) CreateImage() framework.Image {
//	return CreateImage(t)
//}
//
func (t *Theme) CreateLabel() framework.Label {
	return createLabel(t)
}

//
//func (t *Theme) CreateLinearLayout() framework.LinearLayout {
//	return CreateLinearLayout(t)
//}
//
//func (t *Theme) CreateList() framework.List {
//	return CreateList(t)
//}
//
//func (t *Theme) CreatePanelHolder() framework.PanelHolder {
//	return CreatePanelHolder(t)
//}
//
//func (t *Theme) CreateProgressBar() framework.ProgressBar {
//	return CreateProgressBar(t)
//}
//
//func (t *Theme) CreateScrollBar() framework.ScrollBar {
//	return CreateScrollBar(t)
//}
//
//func (t *Theme) CreateScrollLayout() framework.ScrollLayout {
//	return CreateScrollLayout(t)
//}
//
//func (t *Theme) CreateSplitterLayout() framework.SplitterLayout {
//	return CreateSplitterLayout(t)
//}
//
//func (t *Theme) CreateTableLayout() framework.TableLayout {
//	return CreateTableLayout(t)
//}
//
//func (t *Theme) CreateTextBox() framework.TextBox {
//	return CreateTextBox(t)
//}
//
//func (t *Theme) CreateTree() framework.Tree {
//	return CreateTree(t)
//}
//
func (t *Theme) CreateWindow(width, height int, title string) framework.Window {
	return createWindow(t, width, height, title)
}
