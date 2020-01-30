// Copyright 2019 spaGO Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fofe

import (
	"saientist.dev/spago/pkg/mat"
	"saientist.dev/spago/pkg/utils"
)

// The Fixed-Size Ordinally-Forgetting Encoding
//    zt = α * zt−1 + et (1 ≤ t ≤ T)
func Encode(alpha float64, size int, seq []int) []*mat.Sparse {
	var z []*mat.Sparse
	for _, i := range seq {
		if len(z) > 0 {
			t := z[len(z)-1].Clone().(*mat.Sparse)
			t.ProdScalarInPlace(alpha)
			t.AddInPlace(mat.OneHotSparse(size, i))
			z = append(z, t)
		} else {
			z = append(z, mat.OneHotSparse(size, i))
		}
	}
	return z
}

func BiEncode(alpha float64, size int, seq []int) (fwd []*mat.Sparse, bwd []*mat.Sparse) {
	fwd = Encode(alpha, size, seq)
	bwd = Encode(alpha, size, utils.ReverseIntSlice(seq))
	utils.ReverseInPlace(bwd)
	return
}
