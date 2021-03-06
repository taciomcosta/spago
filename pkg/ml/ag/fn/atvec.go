// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import (
	"github.com/nlpodyssey/spago/pkg/mat"
)

var _ Function = &AtVec{}

type AtVec struct {
	x Operand
	i int
}

func NewAtVec(x Operand, i int) *AtVec {
	return &AtVec{x: x, i: i}
}

// Forward computes the output of the function.
func (r *AtVec) Forward() mat.Matrix {
	return mat.NewScalar(r.x.Value().AtVec(r.i))
}

func (r *AtVec) Backward(gy mat.Matrix) {
	if r.x.RequiresGrad() {
		dx := mat.NewEmptyDense(r.x.Value().Dims())
		defer mat.ReleaseDense(dx)
		dx.SetVec(r.i, gy.Scalar())
		r.x.PropagateGrad(dx)
	}
}
