// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package outer

import (
	"github.com/tmacychen/UFG/framework"
)

type Attachable interface {

	// Attached returns true if the control is directly or indirectly attached
	// to a window.
	Attached() bool

	// Attach is called when the control is directly or indirectly attached to a
	// window.
	// Attach should only be called by the parent of the control.
	Attach()

	// Detach is called when the control is directly or indirectly detached from a
	// window.
	// Detach should only be called by the parent of the control.
	Detach()

	// OnAttach subscribes f to be called whenever the control is attached.
	// OnDetach subscribes f to be called whenever the control is detached.
	OnAttach(func()) framework.EventSubscription
	OnDetach(func()) framework.EventSubscription
}
