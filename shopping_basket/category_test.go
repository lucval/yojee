/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "os"
  "reflect"
  "testing"
  "github.com/lucval/yojee/kv"
)

var bookCategory Category = Category{"book", true}
var foodCategory Category = Category{"food", true}
var medicalCategory Category = Category{"medical", true}

// Ensure that categoryMap is correctly populated.
func TestLoadCategoryMap(t *testing.T) {
  // Open KVDB
  kv.Open("category_test.db")

  // Populate category bucket
  kv.Insert("category", "book", "true")
  kv.Insert("category", "food", "true")
  kv.Insert("category", "medical", "true")

  // Populate product bucket
  kv.Insert("product", "book", "book")
  kv.Insert("product", "chocolate bar", "food")
  kv.Insert("product", "box of chocolates", "food")
  kv.Insert("product", "packet of headache pills", "medical")

  // Close KVDB
  kv.Close()

  // Test map
  testCategoryMap := map[string]Category {
    "book":                     bookCategory,
    "chocolate bar":            foodCategory,
    "box of chocolates":        foodCategory,
    "packet of headache pills": medicalCategory,
  }

  // Load map from file
  LoadCategoryMap("category_test.db")

  // Remove KVDB file
  os.Remove("category_test.db")

  for k, v := range categoryMap {
    if reflect.DeepEqual(v, testCategoryMap[k]) == false {
      t.Errorf("Category for product '%s' unmatched (%s instead of %s)",
        k, v, testCategoryMap[k])
    }
  }
}

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
