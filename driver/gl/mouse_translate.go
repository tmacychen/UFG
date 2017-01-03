// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	"fmt"
	"github.com/tmacychen/UFG/framework"

	"github.com/goxjs/glfw"
)

func translateMouseButton(button glfw.MouseButton) framework.MouseButton {
	switch button {
	case glfw.MouseButtonLeft:
		return framework.MouseButtonLeft
	case glfw.MouseButtonMiddle:
		return framework.MouseButtonMiddle
	case glfw.MouseButtonRight:
		return framework.MouseButtonRight
	default:
		panic(fmt.Errorf("Unknown mouse button %v", button))
	}
}

func getMouseState(w *glfw.Window) framework.MouseState {
	var s framework.MouseState
	for _, button := range []glfw.MouseButton{glfw.MouseButtonLeft, glfw.MouseButtonMiddle, glfw.MouseButtonRight} {
		if w.GetMouseButton(button) == glfw.Press {
			s |= 1 << uint(translateMouseButton(button))
		}
	}
	return s
}
