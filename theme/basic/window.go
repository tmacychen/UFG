package basic

import (
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/widget/tools"
	"github.com/tmacychen/UFG/widget"
)

func createWindow(theme *Theme, width, height int, title string) framework.Window {
//	println("createWindow")
	w := &widget.Window{}
	w.Init(w, theme.Driver(), width, height, title)
	w.SetBackgroundBrush(tools.CreateBrush(theme.WindowBackground))
	return w
}