KV
==

*kv* is a library that can be used to provide CRUDL operations on an internal
Bolt KVDB.

A service to execute the same operations from CLI is also implemented.

Getting Started
===============

### Installing

To start using this package, install Go and run `go get`:

```sh
$ go get github.com/lucval/yojee/kv
```

Further install the desired command by running:

```sh
$ go install github.com/lucval/yojee/kv/cmd/kv
```

This will install the requested command line utility into your $GOBIN path.

### Usage

```sh
kv command [arguments]

command:

    get		read key value pair
    set		write key value pair
    del		remove key value pair

Use "kv [command] -h" for more information about a command.

```
