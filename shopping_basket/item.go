/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "fmt"
)

var basicTax float64 = 0.0
var importTax float64 = 0.0
var roundUnit float64 = 0.0

// SetupTaxes sets basic and import taxes to the provided values.
func SetupTaxes(bTax, iTax, rUnit float64) {
  basicTax = bTax
  importTax = iTax
  roundUnit = rUnit
}

type Item struct {
  Quantity  uint64
  Product   *Product
  Price     float64
  Taxes     float64
}

// NewItem instantiates an Item object and calculates the taxes owed over the
// item's product.
// This function returns a pointer to the instantiated Item.
func NewItem(quantity uint64, product *Product, price float64) *Item {
  taxes := 0.0
  if product.Category.Exemption == false {
    // Basic tax on not exempt product
    taxes += roundUp((price * basicTax), roundUnit)
  }
  if product.Imported {
    // Import tax on imported product
    taxes += roundUp((price * importTax), roundUnit)
  }

  return &Item{quantity, product, price, taxes}
}

// String prints Item details.
// This function return a string object.
func (i *Item) String() string {
  return fmt.Sprintf("%d, %s, %.2f", i.Quantity, i.Product, i.Price+i.Taxes)
}
