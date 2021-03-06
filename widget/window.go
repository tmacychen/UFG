package widget

import (
	"github.com/tmacychen/UFG/controller"
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/framework/outer"
	"github.com/tmacychen/UFG/math"
	"github.com/tmacychen/UFG/widget/tools"
)

type WindowOuter interface {
	framework.Window
	outer.Attachable
	outer.IsVisibler
	outer.LayoutChildren
	outer.PaintChilder
	outer.Painter
	outer.Parenter
	outer.Sized
}

type Window struct {
	controller.Container
	controller.Attachable
	controller.Paddable
	controller.PaintChildren
	BackgroundBorderPainter

	outer                 WindowOuter
	driver                framework.Driver
	viewport              framework.Viewport
	viewportSubscriptions []framework.EventSubscription
	windowedSize          math.Size
	focusController       *controller.FocusController
	layoutPending         bool
	drawPending           bool
	updatePending         bool

	onClose       framework.Event // Raised by viewport
	onResize      framework.Event // Raised by viewport
	onMouseMove   framework.Event // Raised by viewport
	onMouseEnter  framework.Event // Raised by viewport
	onMouseExit   framework.Event // Raised by viewport
	onMouseDown   framework.Event // Raised by viewport
	onMouseUp     framework.Event // Raised by viewport
	onMouseScroll framework.Event // Raised by viewport
	onKeyDown     framework.Event // Raised by viewport
	onKeyUp       framework.Event // Raised by viewport
	onKeyRepeat   framework.Event // Raised by viewport
	onKeyStroke   framework.Event // Raised by viewport

	onClick       framework.Event // Raised by MouseController
	onDoubleClick framework.Event // Raised by MouseController
}

func (w *Window) Init(outer WindowOuter, theme framework.Theme, width, height int, title string) {
	//	println("window init")
	w.Container.Init(outer, theme)
	w.Attachable.Init(outer)
	w.Paddable.Init(outer)
	w.PaintChildren.Init(outer)
	w.BackgroundBorderPainter.Init(outer)
	w.outer = outer
	w.driver = theme.Driver()

	w.SetBorderPen(tools.TransparentPen)
	w.setViewport(w.driver.CreateWindowedViewport(width, height, title))

	w.onClose = controller.CreateEvent(func() {})
	w.onResize = controller.CreateEvent(func() {})
	w.onMouseMove = controller.CreateEvent(func(framework.MouseEvent) {})
	w.onMouseEnter = controller.CreateEvent(func(framework.MouseEvent) {})
	w.onMouseExit = controller.CreateEvent(func(framework.MouseEvent) {})
	w.onMouseDown = controller.CreateEvent(func(framework.MouseEvent) {})
	w.onMouseUp = controller.CreateEvent(func(framework.MouseEvent) {})
	w.onMouseScroll = controller.CreateEvent(func(framework.MouseEvent) {})
	w.onKeyDown = controller.CreateEvent(func(framework.KeyboardEvent) {})
	w.onKeyUp = controller.CreateEvent(func(framework.KeyboardEvent) {})
	w.onKeyRepeat = controller.CreateEvent(func(framework.KeyboardEvent) {})
	w.onKeyStroke = controller.CreateEvent(func(framework.KeyStrokeEvent) {})

	w.onClick = controller.CreateEvent(func(framework.MouseEvent) {})
	w.onDoubleClick = controller.CreateEvent(func(framework.MouseEvent) {})

	w.onResize.Listen(func() {
		w.outer.LayoutChildren()
		w.Draw()
	})
	//window 需要显示
	w.Attach()
	//interface compliance test
	_ = framework.Window(w)
	println("window init end")
}

func (w *Window) Title() string {
	return w.viewport.Title()
}

func (w *Window) SetTitle(t string) {
	w.viewport.SetTitle(t)
}
func (w *Window) Scale() float32 {
	return w.viewport.Scale()
}
func (w *Window) SetScale(scale float32) {
	w.viewport.SetScale(scale)
}

func (w *Window) Position() math.Point {
	return w.viewport.Position()
}
func (w *Window) SetPosition(pos math.Point) {
	w.viewport.SetPosition(pos)
}
func (w *Window) Fullscreen() bool {
	return w.viewport.Fullscreen()
}
func (w *Window) SetFullscreen(fullscreen bool) {
	println("window setfullsreen")
	title := w.viewport.Title()
	if fullscreen != w.Fullscreen() {
		old := w.viewport
		if fullscreen {
			// 设置全屏
			w.windowedSize = old.SizeDips()
			w.setViewport(w.driver.CreateFullscreenViewport(0, 0, title))
		} else {
			// 设置为原始大小
			width, height := w.windowedSize.WH()
			w.setViewport(w.driver.CreateWindowedViewport(width, height, title))
		}
		old.Close()
	}
}

func (w *Window) Show() {
	w.Attach()
	w.viewport.Show()
}

func (w *Window) Hide() {
	//w.Detch()
	w.viewport.Hide()
}
func (w *Window) Close() {
	//w.Detch()
	w.viewport.Close()
}

func (w *Window) Focus() framework.Focusable {
	return w.focusController.Focus()

}

func (w *Window) SetFocus(c framework.Control) bool {
	return false
}
func (w *Window) IsVisible() bool {
	return true
}

func (w *Window) requestUpdate() {
	if !w.updatePending {
		w.updatePending = true
		w.driver.Call(w.update)
	}
}

func (w *Window) update() {
	if !w.Attached() {
		// Window was detached between requestUpdate() and update()
		w.updatePending = false
		w.layoutPending = false
		w.drawPending = false
		return
	}
	w.updatePending = false
	if w.layoutPending {
		w.layoutPending = false
		w.drawPending = true
		w.outer.LayoutChildren()
	}
	if w.drawPending {
		w.drawPending = false
		w.Draw()
	}
}
func (w *Window) LayoutChildren() {
	s := w.Size().Contract(w.Padding()).Max(math.ZeroSize)
	o := w.Padding().LT()
	for _, c := range w.outer.Children() {
		c.Layout(c.Control.DesiredSize(math.ZeroSize, s).Rect().Offset(o))
	}

}
func (w *Window) Parent() framework.Parent {
	return nil
}

func (w *Window) Paint(c framework.Canvas) {
	w.PaintBackground(c, c.Size().Rect())
	w.PaintChildren.Paint(c)
	w.PaintBorder(c, c.Size().Rect())
}

func (w *Window) Redraw() {
	w.drawPending = true
	w.requestUpdate()
	println("w.Redraw()")
}
func (w *Window) Relayout() {
	println("w.Relayout()")
	w.layoutPending = true
	w.requestUpdate()
}
func (w *Window) SetSize(size math.Size) {
	w.viewport.SetSizeDips(size)
}
func (w *Window) Size() math.Size {
	return w.viewport.SizeDips()
}

func (w *Window) Draw() framework.Canvas {
	if s := w.viewport.SizeDips(); s != math.ZeroSize {
		c := w.driver.CreateCanvas(s)
		w.outer.Paint(c)
		c.Complete()
		w.viewport.SetCanvas(c)
		return c
	} else {
		return nil
	}
}

func (w *Window) setViewport(v framework.Viewport) {
	for _, s := range w.viewportSubscriptions {
		s.Unlisten()
	}
	w.viewport = v
	w.viewportSubscriptions = []framework.EventSubscription{
		v.OnClose(func() { w.onClose.Fire() }),
		v.OnResize(func() { w.onResize.Fire() }),
		v.OnMouseMove(func(ev framework.MouseEvent) { w.onMouseMove.Fire(ev) }),
		v.OnMouseEnter(func(ev framework.MouseEvent) { w.onMouseEnter.Fire(ev) }),
		v.OnMouseExit(func(ev framework.MouseEvent) { w.onMouseExit.Fire(ev) }),
		v.OnMouseDown(func(ev framework.MouseEvent) { w.onMouseDown.Fire(ev) }),
		v.OnMouseUp(func(ev framework.MouseEvent) { w.onMouseUp.Fire(ev) }),
		v.OnMouseScroll(func(ev framework.MouseEvent) { w.onMouseScroll.Fire(ev) }),
		v.OnKeyDown(func(ev framework.KeyboardEvent) { w.onKeyDown.Fire(ev) }),
		v.OnKeyUp(func(ev framework.KeyboardEvent) { w.onKeyUp.Fire(ev) }),
		v.OnKeyRepeat(func(ev framework.KeyboardEvent) { w.onKeyRepeat.Fire(ev) }),
		v.OnKeyStroke(func(ev framework.KeyStrokeEvent) { w.onKeyStroke.Fire(ev) }),
	}
	w.Relayout()
}

func (w *Window) OnClose(f func()) framework.EventSubscription {
	return w.onClose.Listen(f)
}

func (w *Window) OnResize(f func()) framework.EventSubscription {
	return w.onResize.Listen(f)
}

func (w *Window) OnClick(f func(framework.MouseEvent)) framework.EventSubscription {
	return w.onClick.Listen(f)
}

func (w *Window) OnDoubleClick(f func(framework.MouseEvent)) framework.EventSubscription {
	return w.onDoubleClick.Listen(f)
}

func (w *Window) OnMouseMove(f func(framework.MouseEvent)) framework.EventSubscription {
	return w.onMouseMove.Listen(func(ev framework.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseEnter(f func(framework.MouseEvent)) framework.EventSubscription {
	return w.onMouseEnter.Listen(func(ev framework.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseExit(f func(framework.MouseEvent)) framework.EventSubscription {
	return w.onMouseExit.Listen(func(ev framework.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseDown(f func(framework.MouseEvent)) framework.EventSubscription {
	return w.onMouseDown.Listen(func(ev framework.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseUp(f func(framework.MouseEvent)) framework.EventSubscription {
	return w.onMouseUp.Listen(func(ev framework.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnMouseScroll(f func(framework.MouseEvent)) framework.EventSubscription {
	return w.onMouseScroll.Listen(func(ev framework.MouseEvent) {
		ev.Window = w
		ev.WindowPoint = ev.Point
		f(ev)
	})
}

func (w *Window) OnKeyDown(f func(framework.KeyboardEvent)) framework.EventSubscription {
	return w.onKeyDown.Listen(f)
}

func (w *Window) OnKeyUp(f func(framework.KeyboardEvent)) framework.EventSubscription {
	return w.onKeyUp.Listen(f)
}

func (w *Window) OnKeyRepeat(f func(framework.KeyboardEvent)) framework.EventSubscription {
	return w.onKeyRepeat.Listen(f)
}

func (w *Window) OnKeyStroke(f func(framework.KeyStrokeEvent)) framework.EventSubscription {
	return w.onKeyStroke.Listen(f)
}
