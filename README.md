# Base32
[![Github Status](https://github.com/jpedro/base32/workflows/test/badge.svg)](https://github.com/jpedro/base32/actions)
[![GoDoc](https://godoc.org/github.com/jpedro/base32?status.svg)](https://godoc.org/github.com/jpedro/base32)

The universal base32 cli for MacOS and others.


### Install

```
go install github.com/jpedro/base32
```


### Usage

```bash
# Encode
# You can also use 'enc', '-e', '--encode'
$ base32 test
ORSXG5A=

# Decode
# Can also use 'dec', '-d', '--decode'
$ base32 decode ORSXG5A=
test

# Chain commands
$ echo test | base32 | base32 -d
