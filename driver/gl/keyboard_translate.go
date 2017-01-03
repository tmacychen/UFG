// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gl

import (
	"UFG/framework"

	"github.com/goxjs/glfw"
)

func translateKeyboardKey(in glfw.Key) framework.KeyboardKey {
	switch in {
	case glfw.KeySpace:
		return framework.KeySpace
	case glfw.KeyApostrophe:
		return framework.KeyApostrophe
	case glfw.KeyComma:
		return framework.KeyComma
	case glfw.KeyMinus:
		return framework.KeyMinus
	case glfw.KeyPeriod:
		return framework.KeyPeriod
	case glfw.KeySlash:
		return framework.KeySlash
	case glfw.Key0:
		return framework.Key0
	case glfw.Key1:
		return framework.Key1
	case glfw.Key2:
		return framework.Key2
	case glfw.Key3:
		return framework.Key3
	case glfw.Key4:
		return framework.Key4
	case glfw.Key5:
		return framework.Key5
	case glfw.Key6:
		return framework.Key6
	case glfw.Key7:
		return framework.Key7
	case glfw.Key8:
		return framework.Key8
	case glfw.Key9:
		return framework.Key9
	case glfw.KeySemicolon:
		return framework.KeySemicolon
	case glfw.KeyEqual:
		return framework.KeyEqual
	case glfw.KeyA:
		return framework.KeyA
	case glfw.KeyB:
		return framework.KeyB
	case glfw.KeyC:
		return framework.KeyC
	case glfw.KeyD:
		return framework.KeyD
	case glfw.KeyE:
		return framework.KeyE
	case glfw.KeyF:
		return framework.KeyF
	case glfw.KeyG:
		return framework.KeyG
	case glfw.KeyH:
		return framework.KeyH
	case glfw.KeyI:
		return framework.KeyI
	case glfw.KeyJ:
		return framework.KeyJ
	case glfw.KeyK:
		return framework.KeyK
	case glfw.KeyL:
		return framework.KeyL
	case glfw.KeyM:
		return framework.KeyM
	case glfw.KeyN:
		return framework.KeyN
	case glfw.KeyO:
		return framework.KeyO
	case glfw.KeyP:
		return framework.KeyP
	case glfw.KeyQ:
		return framework.KeyQ
	case glfw.KeyR:
		return framework.KeyR
	case glfw.KeyS:
		return framework.KeyS
	case glfw.KeyT:
		return framework.KeyT
	case glfw.KeyU:
		return framework.KeyU
	case glfw.KeyV:
		return framework.KeyV
	case glfw.KeyW:
		return framework.KeyW
	case glfw.KeyX:
		return framework.KeyX
	case glfw.KeyY:
		return framework.KeyY
	case glfw.KeyZ:
		return framework.KeyZ
	case glfw.KeyLeftBracket:
		return framework.KeyLeftBracket
	case glfw.KeyBackslash:
		return framework.KeyBackslash
	case glfw.KeyRightBracket:
		return framework.KeyRightBracket
	case glfw.KeyGraveAccent:
		return framework.KeyGraveAccent
	case glfw.KeyWorld1:
		return framework.KeyWorld1
	case glfw.KeyWorld2:
		return framework.KeyWorld2
	case glfw.KeyEscape:
		return framework.KeyEscape
	case glfw.KeyEnter:
		return framework.KeyEnter
	case glfw.KeyTab:
		return framework.KeyTab
	case glfw.KeyBackspace:
		return framework.KeyBackspace
	case glfw.KeyInsert:
		return framework.KeyInsert
	case glfw.KeyDelete:
		return framework.KeyDelete
	case glfw.KeyRight:
		return framework.KeyRight
	case glfw.KeyLeft:
		return framework.KeyLeft
	case glfw.KeyDown:
		return framework.KeyDown
	case glfw.KeyUp:
		return framework.KeyUp
	case glfw.KeyPageUp:
		return framework.KeyPageUp
	case glfw.KeyPageDown:
		return framework.KeyPageDown
	case glfw.KeyHome:
		return framework.KeyHome
	case glfw.KeyEnd:
		return framework.KeyEnd
	case glfw.KeyCapsLock:
		return framework.KeyCapsLock
	case glfw.KeyScrollLock:
		return framework.KeyScrollLock
	case glfw.KeyNumLock:
		return framework.KeyNumLock
	case glfw.KeyPrintScreen:
		return framework.KeyPrintScreen
	case glfw.KeyPause:
		return framework.KeyPause
	case glfw.KeyF1:
		return framework.KeyF1
	case glfw.KeyF2:
		return framework.KeyF2
	case glfw.KeyF3:
		return framework.KeyF3
	case glfw.KeyF4:
		return framework.KeyF4
	case glfw.KeyF5:
		return framework.KeyF5
	case glfw.KeyF6:
		return framework.KeyF6
	case glfw.KeyF7:
		return framework.KeyF7
	case glfw.KeyF8:
		return framework.KeyF8
	case glfw.KeyF9:
		return framework.KeyF9
	case glfw.KeyF10:
		return framework.KeyF10
	case glfw.KeyF11:
		return framework.KeyF11
	case glfw.KeyF12:
		return framework.KeyF12
	case glfw.KeyKP0:
		return framework.KeyKp0
	case glfw.KeyKP1:
		return framework.KeyKp1
	case glfw.KeyKP2:
		return framework.KeyKp2
	case glfw.KeyKP3:
		return framework.KeyKp3
	case glfw.KeyKP4:
		return framework.KeyKp4
	case glfw.KeyKP5:
		return framework.KeyKp5
	case glfw.KeyKP6:
		return framework.KeyKp6
	case glfw.KeyKP7:
		return framework.KeyKp7
	case glfw.KeyKP8:
		return framework.KeyKp8
	case glfw.KeyKP9:
		return framework.KeyKp9
	case glfw.KeyKPDecimal:
		return framework.KeyKpDecimal
	case glfw.KeyKPDivide:
		return framework.KeyKpDivide
	case glfw.KeyKPMultiply:
		return framework.KeyKpMultiply
	case glfw.KeyKPSubtract:
		return framework.KeyKpSubtract
	case glfw.KeyKPAdd:
		return framework.KeyKpAdd
	case glfw.KeyKPEnter:
		return framework.KeyKpEnter
	case glfw.KeyKPEqual:
		return framework.KeyKpEqual
	case glfw.KeyLeftShift:
		return framework.KeyLeftShift
	case glfw.KeyLeftControl:
		return framework.KeyLeftControl
	case glfw.KeyLeftAlt:
		return framework.KeyLeftAlt
	case glfw.KeyLeftSuper:
		return framework.KeyLeftSuper
	case glfw.KeyRightShift:
		return framework.KeyRightShift
	case glfw.KeyRightControl:
		return framework.KeyRightControl
	case glfw.KeyRightAlt:
		return framework.KeyRightAlt
	case glfw.KeyRightSuper:
		return framework.KeyRightSuper
	case glfw.KeyMenu:
		return framework.KeyMenu
	default:
		return framework.KeyUnknown
	}
}

func translateKeyboardModifier(in glfw.ModifierKey) framework.KeyboardModifier {
	out := framework.ModNone
	if in&glfw.ModShift != 0 {
		out |= framework.ModShift
	}
	if in&glfw.ModControl != 0 {
		out |= framework.ModControl
	}
	if in&glfw.ModAlt != 0 {
		out |= framework.ModAlt
	}
	if in&glfw.ModSuper != 0 {
		out |= framework.ModSuper
	}
	return out
}
