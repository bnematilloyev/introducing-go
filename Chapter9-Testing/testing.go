package math

import (
	"math/Chapter8-Packages/math"
	"testing"
)

func TestAverage(t *testing.T) {
	v := math.Average([]float64{1, 2})
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
