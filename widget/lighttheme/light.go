// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package lighttheme

import (
	"fmt"
	"UFG/font"
	"UFG/framework"
	"UFG/widget"
	"UFG/widget/tools"
)

func CreateTheme(driver framework.Driver) framework.Theme {
	defaultFont, err := font.CreateFont(font.Default, 12)
	if err == nil {
		defaultFont.LoadGlyphs(32, 126)
	} else {
		fmt.Printf("Warning: Failed to load default font - %v\n", err)
	}

	defaultMonospaceFont, err := font.CreateFont(font.Monospace, 12)
	if err == nil {
		defaultFont.LoadGlyphs(32, 126)
	} else {
		fmt.Printf("Warning: Failed to load default monospace font - %v\n", err)
	}

	scrollBarRailDefaultBg := tools.Black
	scrollBarRailDefaultBg.A = 0.7

	scrollBarRailOverBg := tools.Gray20
	scrollBarRailOverBg.A = 0.7

	neonBlue := tools.ColorFromHex(0xFF5C8CFF)
	focus := tools.ColorFromHex(0xFFC4D6FF)

	return &widget.Theme{
		DriverInfo:               driver,
		DefaultFontInfo:          defaultFont,
		DefaultMonospaceFontInfo: defaultMonospaceFont,
		WindowBackground:         tools.White,

		//                                          fontColor    brushColor   penColor
		BubbleOverlayStyle:        widget.CreateStyle(tools.Gray40, tools.Gray20, tools.Gray40, 1.0),
		ButtonDefaultStyle:        widget.CreateStyle(tools.Gray40, tools.White, tools.Gray40, 1.0),
		ButtonOverStyle:           widget.CreateStyle(tools.Gray40, tools.Gray90, tools.Gray40, 1.0),
		ButtonPressedStyle:        widget.CreateStyle(tools.Gray20, tools.Gray70, tools.Gray30, 1.0),
		CodeSuggestionListStyle:   widget.CreateStyle(tools.Gray40, tools.Gray20, tools.Gray10, 1.0),
		DropDownListDefaultStyle:  widget.CreateStyle(tools.Gray40, tools.White, tools.Gray20, 1.0),
		DropDownListOverStyle:     widget.CreateStyle(tools.Gray40, tools.Gray90, tools.Gray50, 1.0),
		FocusedStyle:              widget.CreateStyle(tools.Gray20, tools.Transparent, focus, 1.0),
		HighlightStyle:            widget.CreateStyle(tools.Gray40, tools.Transparent, neonBlue, 2.0),
		LabelStyle:                widget.CreateStyle(tools.Gray40, tools.Transparent, tools.Transparent, 0.0),
		PanelBackgroundStyle:      widget.CreateStyle(tools.Gray40, tools.White, tools.Gray15, 1.0),
		ScrollBarBarDefaultStyle:  widget.CreateStyle(tools.Gray40, tools.Gray30, tools.Gray40, 1.0),
		ScrollBarBarOverStyle:     widget.CreateStyle(tools.Gray40, tools.Gray50, tools.Gray60, 1.0),
		ScrollBarRailDefaultStyle: widget.CreateStyle(tools.Gray40, scrollBarRailDefaultBg, tools.Transparent, 1.0),
		ScrollBarRailOverStyle:    widget.CreateStyle(tools.Gray40, scrollBarRailOverBg, tools.Gray20, 1.0),
		SplitterBarDefaultStyle:   widget.CreateStyle(tools.Gray40, tools.Gray80, tools.Gray40, 1.0),
		SplitterBarOverStyle:      widget.CreateStyle(tools.Gray40, tools.Gray80, tools.Gray50, 1.0),
		TabActiveHighlightStyle:   widget.CreateStyle(tools.Gray30, neonBlue, neonBlue, 0.0),
		TabDefaultStyle:           widget.CreateStyle(tools.Gray40, tools.White, tools.Gray40, 1.0),
		TabOverStyle:              widget.CreateStyle(tools.Gray30, tools.Gray90, tools.Gray50, 1.0),
		TabPressedStyle:           widget.CreateStyle(tools.Gray20, tools.Gray70, tools.Gray30, 1.0),
		TextBoxDefaultStyle:       widget.CreateStyle(tools.Gray40, tools.White, tools.Gray20, 1.0),
		TextBoxOverStyle:          widget.CreateStyle(tools.Gray40, tools.White, tools.Gray50, 1.0),
	}
}
