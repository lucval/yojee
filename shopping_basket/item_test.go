/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "math"
  "testing"
)

type input struct {
  quantity  uint64
  product   *Product
  price     float64
}

// Ensure that Item is correctly instantiated.
func TestNewItem(t *testing.T) {
  SetupTaxes(0.1, 0.05)

  testMap := map[input]float64{
    input{1, NewProduct("book"), 12.49}:                        0.0,
    input{1, NewProduct("music cd"), 14.99}:                    1.5,
    input{1, NewProduct("chocolate bar"), 0.85}:                0.0,
    input{1, NewProduct("imported box of chocolates"), 10.00}:  0.5,
    input{1, NewProduct("imported bottle of perfume"), 47.50}:  7.15,
    input{1, NewProduct("bottle of perfume"), 18.99}:           1.9,
    input{1, NewProduct("packet of headache pills"), 9.75}:     0.0,
    input{1, NewProduct("imported box of chocolates"), 11.25}:  0.6,
    // Different quantity
    input{2, NewProduct("music cd"), 14.99}:                    1.5,
    input{0, NewProduct("music cd"), 14.99}:                    1.5,
    // Negative price
    input{1, NewProduct("music cd"), -14.99}:                  -1.5,
  }

  for input, res := range testMap {
    item := NewItem(input.quantity, input.product, input.price)
    if math.Abs(item.Taxes - res) > tolerance {
      t.Errorf("Item '%s' has invalid taxes (got %f expected %f)",
        item.Product.Name, item.Taxes, res)
    }
  }
}

// Ensure that Item is correctly printed.
func TestItemString(t *testing.T) {
  SetupTaxes(0.1, 0.05)

  testMap := map[input]string{
    input{1, NewProduct("book"), 12.49}:
      "1, book, 12.49",
    input{1, NewProduct("music cd"), 14.99}:
      "1, music cd, 16.49",
    input{1, NewProduct("chocolate bar"), 0.85}:
      "1, chocolate bar, 0.85",
    input{1, NewProduct("imported bottle of perfume"), 47.50}:
      "1, imported bottle of perfume, 54.65",
    input{1, NewProduct("packet of headache pills"), 9.75}:
      "1, packet of headache pills, 9.75",
    // Different quantity
    input{2, NewProduct("music cd"), 14.99}:
      "2, music cd, 16.49",
    input{0, NewProduct("music cd"), 14.99}:
      "0, music cd, 16.49",
    // Negative price
    input{1, NewProduct("music cd"), -14.99}:
      "1, music cd, -16.49",
  }

  for input, res := range testMap {
    item := NewItem(input.quantity, input.product, input.price)
    if res != item.String() {
      t.Errorf("Item '%s' print mismatch (got '%s' expected '%s')",
        item.Product.Name, res, item)
    }
  }
}
