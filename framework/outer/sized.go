// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package outer

import "github.com/tmacychen/UFG/math"

type Sized interface {
	// Size returns the size of the control. If the control is not attached, then
	// the returned size is undefined.
	Size() math.Size

	// SetSize sets the size of the control to the specified value.
	// SetSize should only be called by the parent of the control during layout.
	SetSize(math.Size)

}
