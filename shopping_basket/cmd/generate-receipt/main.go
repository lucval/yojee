package main

import (
  "encoding/csv"
  "flag"
  "log"
  "os"
  "strconv"
  "strings"
  "github.com/lucval/yojee/shopping_basket"
)

var (
  inputFile string
  basicTax  float64
  importTax float64
  roundUnit float64
)

func init() {
  // Set output flags of standard logger
  log.SetFlags(0)

  flag.StringVar(&inputFile, "file", "", "Shopping basket CSV file path")
  flag.Float64Var(&basicTax, "basicTax", 0.1, "Basic sales tax rate")
  flag.Float64Var(&importTax, "importTax", 0.05, "Import tax rate")
  flag.Float64Var(&roundUnit, "roundUnit", 0.05, "Taxes round unit")
  flag.Parse()

  shopping_basket.SetupTaxes(basicTax, importTax, roundUnit)
}

func main() {
  if inputFile == "" {
    log.Print("Input file is required")
    flag.Usage()
    os.Exit(1)
  }

  // Open input file
  csvFile, err := os.Open(inputFile)
  if err != nil {
    log.Fatalf("Failed to open CSV file: %s", err)
  }
  defer csvFile.Close()

  // Read input file
  lines, err := csv.NewReader(csvFile).ReadAll()
  if err != nil {
    log.Fatalf("Failed to read CSV file: %s", err)
  }

  // Extract header line if present
  if strings.EqualFold(strings.TrimSpace(lines[0][0]), "quantity") {
    lines = lines[1:]
  }

  receipt := shopping_basket.NewReceipt()
  for _, line := range lines {
    // Validate line
    if len(line) < 3 {
      log.Fatal("Invalid item present in basket: missing fields")
    }
    if len(line) > 3 {
      log.Fatal("Invalid item present in basket: too many fields")
    }

    // Parse quantity
    quantity, err := strconv.ParseUint(strings.TrimSpace(line[0]), 0, 32)
    if err != nil {
      log.Fatalf("Invalid quantity value '%s'", line[0])
    }

    // New product
    product := shopping_basket.NewProduct(
      strings.TrimSpace(strings.ToLower(line[1])))

    // Parse price
    price, err := strconv.ParseFloat(strings.TrimSpace(line[2]), 16)
    if err != nil {
      log.Fatalf("Invalid price value '%s'", line[2])
    }

    // Add item to receipt
    receipt.AddItem(shopping_basket.NewItem(quantity, product, price))
  }
  // Print out receipt
  log.Printf("%s", receipt)
}
