# Base32
[![Github Status](https://github.com/jpedro/base32/workflows/test/badge.svg)](https://github.com/jpedro/base32/actions)
[![GoDoc](https://godoc.org/github.com/jpedro/base32?status.svg)](https://godoc.org/github.com/jpedro/base32)

The universal base32 binary.


### Install

```
go install github.com/jpedro/base32
```


### Usage

```bash
# Encode: You can use '-e', '--encode' or just pipe the stdin:
$ echo test | base32
ORSXG5A=

$ base32 -e test
ORSXG5A=

# Decode: Must use '-d', '--decode':
$ base32 --decode ORSXG5A=
test

# Chain commands
$ echo test | base32 | base32 -d
```


### Refresh pkg.go.dev

Check that https://proxy.golang.org/github.com/jpedro/base32/@v/list includes
the latest version or tag.

Use https://proxy.golang.org/github.com/jpedro/base32/@v/v0.1.1.mod to update.

Check https://pkg.go.dev/github.com/jpedro/base32 for the latest.
