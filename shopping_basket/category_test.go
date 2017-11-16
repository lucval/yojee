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
    "bOOk":                     Category{},
    "book of law":              Category{},
    "imported book":            Category{},
    "chocolate bar":            foodCategory,
    "box of chocolate":         Category{},
    "box of chocolates":        foodCategory,
    "box_of_chocolates":        Category{},
    "box of headache pills":    Category{},
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
