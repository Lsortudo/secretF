package service

import (
	"testing"
)

func TestDrawPairs(t *testing.T) {
	participants := []string{"Alice", "Bob", "Charlie", "Diana"}
	pairs, err := DrawPairs(participants)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(pairs) != 2 {
		t.Fatalf("expected 2 pairs, got %d", len(pairs))
	}

	seen := make(map[string]bool)
	for _, pair := range pairs {
		if seen[pair.Pair1] || seen[pair.Pair2] {
			t.Fatalf("duplicate participant in pairs: %v", pair)
		}
		seen[pair.Pair1] = true
		seen[pair.Pair2] = true
	}
}
