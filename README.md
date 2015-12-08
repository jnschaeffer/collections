Introduction
============
[![GoDoc](https://godoc.org/github.com/jnschaeffer/collections?status.svg)](https://godoc.org/github.com/jnschaeffer/collections)

Collections 0.1 is a collection of utility types and functions for performing
common operations on collections in Go.

One of the most frequent (and justified) complaints about Go is the relative
lack of generic types and functions for working with collections, such as
lists. This package attempts to alleviate some of the pain of that.

Collections currently supports:
  * Sorting (using `sort.Sort` through the `Sorter` wrapper type)
  * Zipping: taking two collections and merging them in order
  * Splitting: taking a collection and splitting it by some criteria
  * Shuffling: rearranging a collection in random order

Most interfaces are built from smaller, composable interface types like
`Bounded` (which implements `Len`) and `Swappable` (which implements `Swap`).
A strong focus is kept on compatibility with the Go standard library.

Installation
============

Just `go get github.com/jnschaeffer/collections` or clone the repository into
your `$GOPATH`.

Documentation
=============

All documentation is available in godoc format.

Documentation is also available at [https://godoc.org/github.com/jnschaeffer/collections](https://godoc.org/github.com/jnschaeffer/collections)
