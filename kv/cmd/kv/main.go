package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"github.com/lucval/yojee/kv"
)


var (
	// ErrUsage is returned when a usage message was printed and the process
	// should simply exit with an error.
	ErrUsage = errors.New("usage")

	// ErrUnknownCommand is returned when a CLI command is not specified
	ErrUnknownCommand = errors.New("unknown command")

	// ErrPathRequired is returned when the path to a Bolt database is missing
	ErrPathRequired = errors.New("path required")

	// ErrFileNotFound is returned when a Bolt database does not exist
	ErrFileNotFound = errors.New("file not found")

	// ErrBucketRequired is returned when the bucket name is not specified
	ErrBucketRequired = errors.New("bucket required")

	// ErrBucketNotFound is returned when a bucket does not exist in a database
	ErrBucketNotFound = errors.New("bucket not found")

	// ErrKeyRequired is returned when the key is not specified
	ErrKeyRequired = errors.New("key required")

	// ErrKeyNotFound is returned when a key does not exist in a bucket
	ErrKeyNotFound = errors.New("key not found")

	// ErrInvalidValue is returned when a benchmark reads an unexpected value
	ErrInvalidValue = errors.New("invalid value")
)


// General usage message
var mainUsage string = strings.TrimLeft(`
CLI CRUDL commands for bolt databases.

Usage:

	kv [flags] command [arguments]

Use "kv -h" for more information about flags.

command:

    get		read key value pair
    set		write key value pair
    del		remove key value pair

Use "kv [command] -h" for more information about a command.
`, "\n")

func main() {
    if err := run(); err == ErrUsage {
        os.Exit(2)
    } else if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }
}

func run() error {
	// Set output flags of standard logger
  log.SetFlags(0)

	// Parse flags
	flag.Parse()

	// Require a command at the beginning
	args := flag.Args()

	if len(args) == 0 || strings.HasPrefix(args[0], "-") {
		fmt.Fprintln(os.Stderr, mainUsage)
		return ErrUsage
	}

	// Execute command
	switch args[0] {
	case "help":
		fmt.Fprintln(os.Stderr, mainUsage)
		return ErrUsage
	case "get":
		return getCommand(args[1:]...)
	case "set":
		return setCommand(args[1:]...)
	case "del":
		return delCommand(args[1:]...)
	default:
		return ErrUnknownCommand
	}
}

// 'get' usage message
var getCommandUsage string = strings.TrimLeft(`
Read key value pair(s) from a database.

Usage:

	kv [flags] get db-path bucket [key]

When key is not specified, the whole bucket is returned.
`, "\n")

func getCommand(args ...string) error {
	// Parse flags
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	help := fs.Bool("h", false, "")
	if err := fs.Parse(args); err != nil {
		return err
	} else if *help {
		fmt.Fprintln(os.Stderr, getCommandUsage)
		return ErrUsage
	}

	// Require database path
	path := fs.Arg(0)
	if path == "" {
		return ErrPathRequired
	} else if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// Open database
	kv.Open(path)
	defer kv.Close()

	// Get bucket name and optionally the requested key
	bucket, key := fs.Arg(1), fs.Arg(2)
	if bucket == "" {
		return ErrBucketRequired
	}
	if key != "" {
		// Retrieve the value of the requested key
		entry, err := kv.Get(bucket, key)
		if err != nil {
			log.Printf("%s", err)
			return ErrKeyNotFound
		}
		// Print result to stdout
		fmt.Fprintln(os.Stdout, entry)
	} else {
		// If no key is provided, retrieve all bucket's values
		entries, err := kv.List(bucket)
		if err != nil {
			log.Printf("%s", err)
			return ErrBucketNotFound
		}
		// Print result to stdout
		fmt.Fprintln(os.Stdout, entries)
	}

	return nil
}

// 'set' usage message
var setCommandUsage string = strings.TrimLeft(`
Write a key value pair into a database

Usage:

	kv [flags] set db-path bucket key value
`, "\n")

func setCommand(args ...string) error {
	// Parse flags
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	help := fs.Bool("h", false, "")
	if err := fs.Parse(args); err != nil {
		return err
	} else if *help {
		fmt.Fprintln(os.Stderr, setCommandUsage)
		return ErrUsage
	}

	// Require database path
	path := fs.Arg(0)
	if path == "" {
		return ErrPathRequired
	} else if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// Open database
	kv.Open(path)
	defer kv.Close()

	// Get bucket name and key-value pair
	bucket, key, value := fs.Arg(1), fs.Arg(2), fs.Arg(3)
	if bucket == "" {
		return ErrBucketRequired
	}
	if key == "" {
		return ErrKeyRequired
	}
	// Set key-value pair in bucket
	err := kv.Insert(bucket, key, value)
	if err != nil {
		log.Printf("%s", err)
		return ErrInvalidValue
	}
	// Print success message to stdout
	fmt.Fprintln(os.Stdout, "Success")

	return nil
}

// 'del' usage message
var delCommandUsage string = strings.TrimLeft(`
Remove key value pair(s) from a database

Usage:

	kv [flags] del db-path bucket [key]

When key is not specified, the whole bucket is removed.
`, "\n")

func delCommand(args ...string) error {
	// Parse flags
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	help := fs.Bool("h", false, "")
	if err := fs.Parse(args); err != nil {
		return err
	} else if *help {
		fmt.Fprintln(os.Stderr, delCommandUsage)
		return ErrUsage
	}

	// Require database path
	path := fs.Arg(0)
	if path == "" {
		return ErrPathRequired
	} else if _, err := os.Stat(path); os.IsNotExist(err) {
		return ErrFileNotFound
	}

	// Open database
	kv.Open(path)
	defer kv.Close()

	// Get bucket and optionally the key to be deleted
	bucket, key := fs.Arg(1), fs.Arg(2)
	if bucket == "" {
		return ErrBucketRequired
	}
	var err error
	if key == "" {
		// If not key is provided, drop bucket
		err = kv.RemoveBucket(bucket)
	} else {
		// Remove key from bucket
		err = kv.Remove(bucket, key)
	}
	if err != nil {
		log.Printf("%s", err)
		return ErrInvalidValue
	}
	// Print success message to stdout
	fmt.Fprintln(os.Stdout, "Success")

	return nil
}
