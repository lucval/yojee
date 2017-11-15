/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "bytes"
  "fmt"
)

type Receipt struct {
  Items   []*Item
  Taxes   float64
  Total   float64
}

// NewReceipt instantiates an empty Receipt object.
// This function returns a pointer to the instantiated Receipt.
func NewReceipt() *Receipt {
  return &Receipt{}
}

// AddItem appends Items to Receipt and increments the receipt's taxes and
// total accordingly.
func (r *Receipt) AddItem(item *Item) {
  r.Items = append(r.Items, item)
  r.Taxes += float64(item.Quantity) * item.Taxes
  r.Total += float64(item.Quantity) * (item.Price + item.Taxes)
}

// String pretty prints Receipt details.
// This function return a string object.
func (r *Receipt) String() string {
  var buffer bytes.Buffer

  for _, item := range r.Items {
    buffer.WriteString(item.String())
    buffer.WriteString("\n")
  }

  buffer.WriteString(fmt.Sprintf("\nSales Taxes: %.2f\n", r.Taxes))
  buffer.WriteString(fmt.Sprintf("Total: %.2f", r.Total))

  return buffer.String()
}
