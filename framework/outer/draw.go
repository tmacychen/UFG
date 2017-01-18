// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package outer

import (
	"github.com/tmacychen/UFG/framework"
)

type Draw interface {
	// Draw draws the control's visual apperance into the returned, new canvas.
	// Draw is typically called by the parent of the control - calling Draw will
	// not issue a re-draw of an attached control.
	Draw() framework.Canvas
}
