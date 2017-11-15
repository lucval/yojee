/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "reflect"
  "testing"
)

// Ensure that Category is correctly instantiated.
func TestNewCategory(t *testing.T) {
  testMap := map[string]Category{
    "book":                     bookCategory,
    "bOOk":                     otherCategory,
    "book of law":              otherCategory,
    "imported book":            otherCategory,
    "chocolate bar":            foodCategory,
    "box of chocolate":         otherCategory,
    "box of chocolates":        foodCategory,
    "box_of_chocolates":        otherCategory,
    "box of headache pills":    otherCategory,
    "packet of headache pills": medicalCategory,
  }

  for pName, cat := range testMap {
    c := NewCategory(pName)
    if reflect.DeepEqual(*c, cat) == false {
      t.Errorf("Category for product '%s' unmatched (%s instead of %s)",
        pName, cat.Name, c.Name)
    }
  }
}
