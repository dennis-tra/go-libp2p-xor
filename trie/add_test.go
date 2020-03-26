package trie

import (
	"math/rand"
	"testing"

	"github.com/libp2p/go-libp2p-xor/key"
)

// Verify mutable and immutable add do the same thing.
func TestMutableAndImmutableAddSame(t *testing.T) {
	for _, s := range testAddSameSamples {
		mut := New()
		immut := New()
		for _, k := range s.Keys {
			mut.Add(k)
			immut = Add(immut, k)
		}
		if !Equal(mut, immut) {
			t.Errorf("mutable trie %v differs from immutable trie %v", mut, immut)
		}
	}
}

func TestAddIsOrderIndependent(t *testing.T) {
	for _, s := range testAddSameSamples {
		base := New()
		for _, k := range s.Keys {
			base.Add(k)
		}
		for j := 0; j < 100; j++ {
			perm := rand.Perm(len(s.Keys))
			reordered := New()
			for i := range s.Keys {
				reordered.Add(s.Keys[perm[i]])
			}
			if !Equal(base, reordered) {
				t.Errorf("trie %v differs from trie %v", base, reordered)
			}
		}
	}
}

type testAddSameSample struct {
	Keys []key.Key
}

var testAddSameSamples = []*testAddSameSample{
	{Keys: []key.Key{{1, 3, 5, 7, 11, 13}}},
	{Keys: []key.Key{{11, 22, 23, 25, 27, 28, 31, 32, 33}}},
}
