# Base32

The universal base32 cli for MacOS and others.


### Install

```
go install github.com/jpedro/base32
```


### Usage

```bash
# Encode
# You can also use 'enc', '-e', '--encode'
$ base32 encode test
ORSXG5A=

# Decode
# Can also use 'dec', '-d', '--decode'
$ base32 decode ORSXG5A=
test

# Chain commands
$ echo test | base32 --encode | base32 -d
