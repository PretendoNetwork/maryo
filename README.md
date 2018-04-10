# maryo
## a proxy server used for accessing pretendo, the open source nintendo network replacement
## *(also known as the terminal app your grandma can use)*

## about
maryo is a proxy program intended to be as easy as possible to use, the user interface comes first, then the features. if i were to focus more on the features, i could end up with a program that works, but i'd then have a bunch of people asking me how to use it.

## instructions

### using precompiled binaries

go to the releases, and download the correct binary for your system. just double-click it and follow the instructions (*it's that easy!*)

### building from source ~~(WARNING: EXTREMELY HARDCORE)~~

#### prerequisites

- [golang](https://golang.org/)
- [goproxy](https://github.com/elazarl/goproxy)
- [ansicolor](https://github.com/shiena/ansicolor)
- [httpscerts](https://github.com/kabukky/httpscerts)

#### building

1. clone this repository
2. it's easy, just `go build` in the source directory
3. you can then double click your own binaries instead of ours (i don't care)

