/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "reflect"
  "testing"
)

// Ensure that imported Product is correctly detected.
func TestIsImported(t *testing.T) {
  type res struct {
    Imported  bool
    Name      string
  }

  testMap := map[string]res{
    "book":                     res{false, "book"},
    "imported book":            res{true, "book"},
    "ImpOrted book":            res{false, "ImpOrted book"},
    "importedbook":             res{false, "importedbook"},
    "book imported":            res{false, "book imported"},
    "imported ":                res{true, ""},
    "imported imported book":   res{true, "imported book"},
  }

  for pName, r := range testMap {
    imported, newName := isImported(pName)
    if (r.Name != newName || r.Imported != imported) {
      t.Errorf("Product '%s' imported detection failed", pName)
    }
  }
}

// Ensure that Product is correctly instantiated.
func TestNewProduct(t *testing.T) {
  testMap := map[string]Product{
    "book":                     Product{"book", &bookCategory, false},
    "bOOk":                     Product{"bOOk", &otherCategory, false},
    "imported book":            Product{"book", &bookCategory, true},
    "chocolate bar":            Product{"chocolate bar", &foodCategory, false},
    "imported chocolate bars":  Product{"chocolate bars", &otherCategory, true},
    "packet of headache pills": Product{
      "packet of headache pills", &medicalCategory, false},
  }

  for pName, prod := range testMap {
    p := NewProduct(pName)
    if reflect.DeepEqual(*p, prod) == false {
      t.Errorf("Product '%s' mismatch (got '%s' expected '%s')", pName, prod, p)
    }
  }
}

// Ensure that Product is correctly printed.
func TestProductString(t *testing.T) {
  testSlice := []string{
    "book", "bOOk", "imported book", "chocolate bar", "imported chocolate bars",
    "packet of headache pills", "packet of headache pills"}

  for _, pName := range testSlice {
    p := NewProduct(pName)
    if pName != p.String() {
      t.Errorf("Product '%s' print mismatch (got '%s')", pName, p)
    }
  }
}
