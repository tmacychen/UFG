package main

import (
	"github.com/tmacychen/UFG/driver/gl"
	"github.com/tmacychen/UFG/framework"
	"github.com/tmacychen/UFG/theme/light"
	"github.com/tmacychen/UFG/font"
	//"github.com/tmacychen/UFG/math"
	"github.com/tmacychen/UFG/widget/tools"
	//"time"
)


func appMain(driver framework.Driver) {
	f, err := font.CreateFont(font.Default, 100)
	if err != nil {
		panic(err)
	}
	println("app main")
	theme := light.CreateTheme(driver)
	window := theme.CreateWindow( 500, 200, "hi")
//	window.SetBackgroundBrush(tools.CreateBrush(tools.Gray70))


	label := theme.CreateLabel()
	label.SetText("Hello world")
	label.SetColor(tools.Blue)
	label.SetFont(f)
	window.AddChild(label)
	//ticker := time.NewTicker(time.Millisecond * 30)
	//go func() {
	//	phase := float32(0)
	//	for _ = range ticker.C {
	//		c := tools.Color{
	//			R: 0.75 + 0.25*math.Cosf((phase+0.000)*math.TwoPi),
	//			G: 0.75 + 0.25*math.Cosf((phase+0.333)*math.TwoPi),
	//			B: 0.75 + 0.25*math.Cosf((phase+0.666)*math.TwoPi),
	//			A: 0.50 + 0.50*math.Cosf(phase*10),
	//		}
	//		phase += 0.01
	//		driver.Call(func() {
	//			label.SetColor(c)
	//		})
	//	}
	//}()

	//window.OnClose(ticker.Stop)
	window.OnClose(driver.Terminate)

}

func main() {
	gl.StartDriver(appMain)
}
