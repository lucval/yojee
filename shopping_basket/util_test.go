/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "math"
  "testing"
)

// Ensure that decimal is correctly rounded up to the nearest unit.
func TestRoundUp(t *testing.T) {
  type input struct {
    f     float64
    unit  float64
  }

  testMap := map[input]float64{
    // Test classic
    input{1.63, 0.05}:    1.65,
    // Test 1 digit fractional
    input{3.8, 0.05}:     3.8,
    // Test 3 digits fractional
    input{0.351, 0.05}:   0.4,
    // Test 4 digits fractional (and new unit)
    input{0.4513, 0.07}:  0.49,
    // Test negative
    input{-1.63, 0.05}:  -1.65,
  }

  for input, res := range testMap {
    r := roundUp(input.f, input.unit)
    if math.Abs(r - res) > tolerance {
      t.Errorf("Failed to round up %f (got %f expected %f)", input.f, r, res)
    }
  }
}
