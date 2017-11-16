Shopping Basket
===============

*shopping_basket* is a package which contains logic related to various
components of a traditional shopping basket.

The base element of a shopping basket is a *product*. Every *product* belongs
to a *category* which determines whether the *product* is exempt from basic
sales tax.
Categories and their mappings with existing products must be provided via a
[Bolt](https://github.com/boltdb/bolt) key/value database (if no category is
found for a specific product, the product will not be exempt from basic sales
tax).
Furthermore in the current implementation *product* names are identifiers used
to map a *product* with the corresponding *category*.
A second import tax can be applied on imported products. This occurs when a
*product* name is prefixed by the word "imported" (case insensitive followed by
a space).

An *item* is intended as an entry (line) of a shopping *receipt* and can
therefore be seen as a collection of same products (but a *product* and thus
items can be repeated multiple times in a *receipt*). Taxes are calculated on
each *item*'s *product* price rounding them up to the nearest round unit.
As already said a *receipt* is a collection of items with the corresponding
taxes and total price.

Getting Started
===============

### Installing

To start using this package, install Go and run `go get`:
```sh
$ go get github.com/lucval/yojee/shopping_basket
```

The [kv library](https://github.com/lucval/yojee/tree/category-from-kv/kv#kv)
must also be retrieved:
```sh
$ go get github.com/lucval/yojee/kv
```

And installed:
```sh
$ go install github.com/lucval/yojee/kv/cmd/kv
```

**After having checked out to the category-from-kv branch**, you can install the desired command by running:
```sh
$ go install github.com/lucval/yojee/shopping_basket/cmd/<command>
```

This will install the requested command line utility into your $GOBIN path.

### Setup database

Categories can be added to the database using this command:
```sh
$ kv set <db-path> category <category-name> <exemption>
```
Where:
- db-path is the path to the Bolt kvdb file (if not yet available a new file
  will be created)
- category-name is the name of the category to be created
- exemption is a boolean representing whether products belonging to this
  category are exempt from basic sales tax

Products can be added to the database using this command:
```sh
$ kv set <db-path> product <product-name> <category-name>
```
Where:
- db-path is the path to the Bolt kvdb file (if not yet available a new file
  will be created)
- product-name is the name of the product (imported prefix is not required)
- category-name is the name of the category to be created

An example key/value database file *category.db* is provided in the repo as well to speed up this procedure.

Commands
========

At the moment the following commands have been implemented.

generate-receipt
----------------
Prints out the receipt details of a shopping basket provided as input in a
CSV file.

### Usage
```sh
generate-receipt [ARGS]

ARGS:
  -csvFile string
    Shopping basket CSV file path
  -kvdbFile string
    Categories KVDB file path
  -basicTax float
    Basic sales tax rate (default 0.1)
  -importTax float
    Import tax rate (default 0.05)
  -roundUnit float
    Taxes round unit (default 0.05)
```

### Input

Shopping baskets must be provided as input in a CSV file as follows:
```sh
quantity, product, price
```
where:
- quantity is a positive integer
- product is the product name
- price is a decimal

Future Improvements
===================
- an extra input field could be used to define imported products
