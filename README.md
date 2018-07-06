# hex2ascii

A command line utility to convert a hex string to ASCII.

## Install

    go get github.com/scottjbarr/hex2ascii

## Usage

Hex values via stdin

```
$ echo 68656c6c6f | hex2ascii
hello

$ echo 68 65 6c 6c 6f | hex2ascii
hello
```

Or hex value as argument

```
$ hex2ascii 68 65 6c 6c 6f
hello
```