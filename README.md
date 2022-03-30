# crank
[![Travis CI](https://img.shields.io/travis/gesquive/crank/master.svg?style=flat-square)](https://travis-ci.org/gesquive/crank)
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/gesquive/crank/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/gesquive/crank)

A simple password generation cli app. This app was originally built as a playground for some ideas.

It can generate passwords in multiple formats:

| random                   | trigraph  | xkcd                            |
|--------------------------|-----------|---------------------------------|
| Z#EZssS6Skyh%xQhP@datwyW | rationatc | find factory bean waste         |
| w#M27Rpx2G7KI50Oc1yCluh5 | ishovenfi | depth team thumb positive       |
| ta5$WoeqdmEyIWgLGL%UwB5d | pedistans | excellent chicken glass ability |
| kQ3j$lfIwSiP3Ks5lSVML%Y$ | litinglen | stairs live thing die           |
| Qr%A7UHdIddMR9K5QskRmypT | nizzinama | draw yard fifty soap            |


## Installing

### Compile
This project requires go1.6+ to compile. Just run `go get -u github.com/gesquive/crank` and the executable should be built for you automatically in your `$GOPATH`.

## Usage
```console
Generate secure random passwords in a number of different formats

Usage:
  crank [flags] [command]

Available Commands:
  random      Generate random passwords
  trigraph    Generate human readable trigraph passwords
  xkcd        Generate passwords with the XKCD password scheme

Flags:
  -c, --config string   config file (default is $HOME/.config/crank.yml)
  -n, --number value    The number of passwords to generate (default 5)
  -v, --version         Show the version and exit

Use "crank [command] --help" for more information about a command.
```

## Documentation

This documentation can be found at github.com/gesquive/crank

## License

This package is made available under an MIT-style license. See LICENSE.

## Credits

Trigraph generation code was largely based on the work of multicians. Their excellent gpw app source code can be found at
http://www.multicians.org/thvv/tvvtools.html

The XKCD password scheme originated here: http://xkcd.com/936/

## Contributing

PRs are always welcome!
