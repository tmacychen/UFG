// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package widget

import (
	"UFG/framework"
)

type ParentableOuter interface{}

type Parentable struct {
	outer  ParentableOuter
	parent framework.Parent
}

func (p *Parentable) Init(outer ParentableOuter) {
	p.outer = outer
}

func (p *Parentable) Parent() framework.Parent {
	return p.parent
}

func (p *Parentable) SetParent(parent framework.Parent) {
	p.parent = parent
}
