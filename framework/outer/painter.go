// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package outer

import (
	"github.com/tmacychen/UFG/framework"
)

type Painter interface {
	Paint(framework.Canvas)
}
