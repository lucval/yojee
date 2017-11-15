/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "bytes" 
  "strings"
)

type Product struct {
  Name      string
  Category  *Category
  Imported  bool
}

// IsImported check whether a product is imported.
// This is true when the product name is prefixed by the keyword "imported"
// followed by a trailing space.
// This function returns a boolean describing whether the product is imported
// and a copy of 'productName' stripped of the "imported" keyword and trailing
// space.
func isImported(productName string) (bool, string) {
  if strings.HasPrefix(productName, "imported ") {
    return true, productName[9:]
  }
  return false, productName
}

// NewProduct instantiates a Product object.
// This function returns a pointer to the instantiated Product.
func NewProduct(productName string) *Product {
  imported, newProductName := isImported(productName)

  return &Product{
    Name:       newProductName,
    Category:   NewCategory(newProductName),
    Imported:   imported,
  }
}

// String prints Product details.
// This function return a string object.
func (p *Product) String() string {
  var buffer bytes.Buffer

  if p.Imported {
    buffer.WriteString("imported ")
  }

  buffer.WriteString(p.Name)

  return buffer.String()
}
