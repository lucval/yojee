/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket


type Category struct {
  Name      string
  Exemption bool
}

var bookCategory Category = Category{"Book", true}
var foodCategory Category = Category{"Food", true}
var medicalCategory Category = Category{"Medical", true}

var categoryMap = map[string]Category {
  "book":                     bookCategory,
  "chocolate bar":            foodCategory,
  "box of chocolates":        foodCategory,
  "packet of headache pills": medicalCategory,
}

// NewCategory instantiates a Category object from a pre-loaded map of
// categories.
// This function returns a pointer to the instantiated Category.
func NewCategory(productName string) *Category {
  c := categoryMap[productName]
  return &c
}
