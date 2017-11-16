/*
Package shopping_basket defines general purposes functions and objects to be
used by several applications.
*/
package shopping_basket

import (
  "log"
  "strconv"
  "github.com/lucval/yojee/kv"
)

type Category struct {
  Name      string
  Exemption bool
}

var categoryMap map[string]Category

// LoadCategoryMap populates the map of categories from a KVDB provided as
// input.
func LoadCategoryMap(dbName string) {
  // Open database
  kv.Open(dbName)

  // Retrieve categories from KVDB
	rawCategories, err := kv.List("category")
	if err != nil {
		log.Fatalf("Failed to lookup categories, please check the KVDB file")
	}

  // Load categories in a map
  categories := make(map[string]Category)
	for k, v := range rawCategories {
    exemption, err := strconv.ParseBool(v)
    if err != nil {
      log.Printf("%s", err)
      log.Fatalf("Tried to load invalid category")
    }
    categories[k] = Category{k, exemption}
  }

  // Retrieve products from KVDB
	rawCategoryMap, err := kv.List("product")
	if err != nil {
		log.Printf("%s", err)
		log.Fatalf("Failed to lookup products")
	}

  // Load categoryMap
  categoryMap = make(map[string]Category)
	for k, v := range rawCategoryMap {
    c := categories[v]
    if c.Name != "" {
      categoryMap[k] = c
    }
  }
}

// NewCategory instantiates a Category object from a pre-loaded map of
// categories.
// This function returns a pointer to the instantiated Category.
func NewCategory(productName string) *Category {
  c := categoryMap[productName]
  if c.Name == "" {
    // Default to category Other
    c.Name = "Other"
    c.Exemption = false
  }
  return &c
}
