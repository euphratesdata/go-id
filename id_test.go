// Copyright 2016, 2017 Marc Wilson, Scorpion Compute. All rights
// reserved. Use of this source code is governed by a
// BSD-style license that can be found in the LICENSE file.

package id

import (
	"testing"
	"math/rand"
)

const seed = 0x12345
const max = 10000

func TestID(test *testing.T) {
	encoder, _ := Encode(16, 10, 36)
	decoder, _ := Decode(16, 10, 36)

	// Test the example found in article.
	ID := uint64(241294492511762325)
	s, t, l := decoder(ID)
	ID_recomputed := encoder(s, t, l)
	if ID != ID_recomputed {
		test.Errorf("encoder(%d, %d, %d) == %d, want %d\n", s, t, l, ID_recomputed, ID)
	}

	// Test random 62-bit ids
	rnd := rand.New(rand.NewSource(seed))
	for i := 1; i <= max; i++ {
		ID := uint64(rnd.Int63n(4611686018427387904))
		s, t, l := decoder(ID)
		ID_recomputed := encoder(s, t, l)
		if ID != ID_recomputed {
			test.Errorf("encoder(%d, %d, %d) == %d, want %d\n", s, t, l, ID_recomputed, ID)
		}
	}
}
