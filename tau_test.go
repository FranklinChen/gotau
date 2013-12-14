package gotau

import (
	"math"
	"testing"
)

func TestTau(t *testing.T) {
	if half_tau := Tau / 2; half_tau != math.Pi {
		t.Errorf("tau/2 = %v, want %v", half_tau, math.Pi)
	}
}

func Testτ(t *testing.T) {
	if half_τ := τ / 2; half_τ != math.Pi {
		t.Errorf("τ/2 = %v, want %v", half_τ, math.Pi)
	}
}
