// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fn

import "github.com/nlpodyssey/spago/pkg/mat"

var _ Function = &ColView{}

type ColView struct {
	x Operand
	i int
}

// Extract the i-th column from the input matrix
func NewColView(x Operand, i int) *ColView {
	if i < 0 {
		panic("fn: invalid column index")
	}
	return &ColView{x: x, i: i}
}

// Forward computes the output of the function.
func (r *ColView) Forward() mat.Matrix {
	xv := r.x.Value()
	rows, cols := xv.Dims()
	if r.i >= cols {
		panic("fn: matrix with not compatible size")
	}
	y := mat.GetDenseWorkspace(1, rows)
	for i := 0; i < rows; i++ {
		y.Set(0, i, xv.At(i, r.i))
	}
	return y
}

func (r *ColView) Backward(gy mat.Matrix) {
	if !(r.x.Value().Rows() == gy.Size()) {
		panic("fn: matrices with not compatible size")
	}
	if r.x.RequiresGrad() {
		gx := mat.NewEmptyDense(r.x.Value().Dims())
		defer mat.ReleaseDense(gx)
		for i := 0; i < r.x.Value().Rows(); i++ {
			gx.Set(i, r.i, gy.At(0, i))
		}
		r.x.PropagateGrad(gx)
	}
}
