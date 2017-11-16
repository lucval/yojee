Shopping Basket
===============

*shopping_basket* is a package which contains logic related to various
components of a traditional shopping basket.

The base element of a shopping basket is a *product*. Every *product* belongs
to a *category* which determines whether the *product* is exempt from basic
sales tax.
In the current implementation 4 categories are defined: book, food, medical and
other. The latter is the only *category* not exempt from basic sales tax. For
convenience the definition of these categories is hard-coded in the source.
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

Further install the desired command by running:

```sh
$ go install github.com/lucval/yojee/shopping_basket/cmd/<command>
```

This will install the requested command line utility into your $GOBIN path.

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
  -file string
    Shopping basket CSV file path
  -basicTax float
    Basic sales tax rate (default 0.1)
  -importTax float
    Import tax rate (default 0.05)
  -roundUnit float
    Taxes round unit (default 0.05)
```

### Input

Shopping baskets must be provided as input in a CSV file as follows.
```sh
quantity, product, price
```
where:
- quantity is a positive integer
- product is the product name
- price is a decimal

Future Improvements
===================
- *product* should be provided with an ID and mapping between products and
categories could be defined by mean of a relational or key/value database
(a version using a pre-populated kvdb is available in a [separate branch](https://github.com/lucval/yojee/tree/category-from-kv/shopping_basket#shopping-basket))
- a relation or an extra input field could be used to define imported products
