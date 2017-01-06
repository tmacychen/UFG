package main

import (
	"github.com/tmacychen/UFG/driver/gl"
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/math"
	"github.com/tmacychen/UFG/theme/light"
)

func appMain(driver framework.Driver) {
	theme := light.CreateTheme(driver)
	win := theme.CreateWindow(200, 150, "Window")
	win.SetScale(1.0)
	win.SetPadding(math.Spacing{L: 10, R: 10, T: 10, B: 10})
	win.OnClose(driver.Terminate)

	button := theme.CreateButton()
	button.SetHorizontalAlignment(framework.AlignCenter)
	button.SetSizeMode(framework.Fill)
	button.SetText("Make fullscreen")
	win.AddChild(button)

}

func main() {
	gl.StartDriver(appMain)
}
