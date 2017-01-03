package main

import (
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/driver/gl"
	"github.com/tmacychen/UFG/theme/light"
	"github.com/tmacychen/UFG/math"
)


func appMain(driver framework.Driver){
	theme := light.CreateTheme(driver)
	win := theme.CreateWindow(200,150,"Window")
	win.OnClose(driver.Terminate)
	win.SetScale(1.0)
	win.SetPadding(math.Spacing{L:10,R:10,T:10,B:10})

}


func main(){
		gl.StartDriver(appMain)
}
