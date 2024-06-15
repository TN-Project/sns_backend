package random

import (
	"testing"
)

func TestGenerateRandomInt(t *testing.T) {
	got := GenerateRandomInt()
	if got < 1000000000 || got > 9999999999 {
		t.Errorf("GenerateRandomInt() = %d; want 1000000000 <= value <= 9999999999", got)
	}
}
