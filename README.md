# HEX

## Description

This software reads file names from the command line
and prints their content as bytes represented hexadecimally.

## Building

On unix-like systems.

```sh
cd ./src
go build -o ../bin/hex .
```

On windows systems.

```cmd
build
```

## Usage

Example 1:

Read `myfile.jpg` and print its bytes as hex values.

```sh
hex myfile.jpg
```

Example 2:

Check out the help.

```sh
hex -h
```
