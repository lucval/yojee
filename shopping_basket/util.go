/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "math"
)

// Tolerance to cope with float loss of precision
const tolerance = 0.00001

// RoundUp rounds up a given decimal 'f' to the nearest 'unit'.
// This function returns a rounded decimal number.
func roundUp(f, unit float64) float64 {
  if f > 0 {
    return math.Ceil(f * (1/unit)) / (1/unit)
  }
  if f < 0 {
    return math.Floor(f * (1/unit)) / (1/unit)
  }
  return f
}
