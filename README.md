# DIF Presentation Exchange

A Golang implementation of the [DIF Presentation Exchange](https://identity.foundation/presentation-exchange). Initially contributed by the team at [Workday Credentials](https://github.com/workdaycredentials/).

## Go
This library uses Go version [1.13](https://golang.org/doc/go1.13).

## Mage
This library uses the [Mage](https://magefile.org/) build tool.

```
$ mage
Targets:
  build         builds the library.
  clean         deletes any build artifacts.
  packr         generates go files for static resources.
  packrClean    deletes all the packr generated go files.
  test          runs unit tests without coverage.
```

## Packr

[Packr](https://github.com/gobuffalo/packr) - The simple and easy way to embed static files into Go binaries.

Golang binaries only contain content from *.go* files. Therefore, static resources like JSON Schemas need to be
converted. The ledger-common code uses the packr library to abstract away the filesystem when accessing static
resource files.  The Packr tool will generate Golang files containing the contents of those
resources. When a developer adds a new static resource file, they must call `mage packr` in order to generate
new *.go* files. These generated files should be committed to source control in order to remove any dependency on
the Packr tool in consumers of this library.  As a convenience, Packr will automatically be installed using
Gobin.

The `mage packrclean` command will delete all existing generated files.
