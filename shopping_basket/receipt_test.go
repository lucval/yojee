/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "math"
  "testing"
)

// Ensure that Item is correctly added to Receipt and taxes and total are
// calculated correctly.
func TestReceiptAddItem(t *testing.T) {
  SetupTaxes(0.1, 0.05)

  testSlice := []Item{
    Item{1, NewProduct("book"), 12.49, 0.0},
    Item{1, NewProduct("music cd"), 14.99, 1.5},
    Item{1, NewProduct("chocolate bar"), 0.85, 0.0},
    Item{1, NewProduct("imported box of chocolates"), 10.00, 0.5},
    Item{1, NewProduct("imported bottle of perfume"), 47.50, 7.15},
    Item{1, NewProduct("bottle of perfume"), 18.99, 1.9},
    Item{1, NewProduct("packet of headache pills"), 9.75, 0.0},
    Item{1, NewProduct("imported box of chocolates"), 11.25, 0.6},
    // Different quantity
    Item{2, NewProduct("music cd"), 14.99, 1.5},
    Item{0, NewProduct("music cd"), 14.99, 1.5},
    // Negative price
    Item{1, NewProduct("music cd"), -14.99, -1.5},
  }

  // Result
  taxes := 13.15
  total := 153.96

  receipt := NewReceipt()
  for _, item := range testSlice {
    receipt.AddItem(&item)
  }
  if math.Abs(receipt.Taxes - taxes) > tolerance {
    t.Errorf("Error in receipt taxes (got %f expected %f)",
      receipt.Taxes, taxes)
  }
  if math.Abs(receipt.Total - total) > tolerance {
    t.Errorf("Error in receipt total (got %f expected %f)",
      receipt.Taxes, taxes)
  }
}

// Ensure that Receipt is correctly printed.
func TestReceiptString(t *testing.T) {
  SetupTaxes(0.1, 0.05)

  receipt := NewReceipt()
  receipt.AddItem(NewItem(1, NewProduct("book"), 12.49))
  receipt.AddItem(NewItem(1, NewProduct("music cd"), 14.99))
  receipt.AddItem(NewItem(1, NewProduct("chocolate bar"), 0.85))

  res := `1, book, 12.49
1, music cd, 16.49
1, chocolate bar, 0.85

Sales Taxes: 1.50
Total: 29.83`

  if res != receipt.String() {
    t.Errorf("Receipt print mismatch\nGot:\n%s\n\nExpected:\n%s", receipt, res)
  }
}
