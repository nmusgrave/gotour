# Links

[Language Documentation](https://golang.org/doc/effective_go.html)

[Tutorial](https://tour.golang.org/welcome/1)

[Project Setup](https://golang.org/doc/code.html)

# Basics

## Compiling, Building, and Running

### Installing

Compiles and installs executable in bin/

Note: executables are statically linked.

```
$ go install gotour
```

or, if inside package, can just do:

```
$ go install
```

### Building

Compile without installing

```
$ go build stringutil
```

or if you're in the directory, just

```
$ go build
```

### Running

```
$ $GOPATH/bin/test
```

or, since the path to bin/ is set in PATH, can just do the following, from anywhere:

```
$ gotour
```

## Coding

### Runes

A 'code point' that represents a single character. Also an alias for 'int32'. A 'rune constant' and 'char constant' are analagous in Go.

### Package Names

Must be present at start of every file:

```
package name
```

with all files in package using same name. Conventionally, is last element of path. Executables always use the 'main' package.

### Testing

Test files are identified by "somename_test.go".  Test functions are named TestSomething. Failures are noticed when `t.Fail` or `t.Error` are invoked.

```
$ go test stringutil
```

or just the following, if in the package directory

```
$ go test
```

### Remote Packages

Import the URL at the head of the file, and execute

```
$ go get github.com/user/repo/package
$ $GOPATH/bin/package
```

