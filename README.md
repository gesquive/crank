# crank
[![Software License](https://img.shields.io/badge/License-MIT-orange.svg?style=flat-square)](https://github.com/gesquive/crank/blob/master/LICENSE)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/gesquive/crank)
[![Build Status](https://img.shields.io/circleci/build/github/gesquive/crank?style=flat-square)](https://circleci.com/gh/gesquive/crank)
[![Coverage Report](https://img.shields.io/codecov/c/gh/gesquive/crank?style=flat-square)](https://codecov.io/gh/gesquive/crank)

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
This project requires go1.17+ to compile. Just run `go install github.com/gesquive/crank@latest` and the executable should be built for you automatically in your `$GOPATH`.

Optionally you can run `make install` to build and copy the executable to `/usr/local/bin/` with correct permissions.

### Download
Alternately, you can download the latest release for your platform from [github](https://github.com/gesquive/crank/releases/latest).

Once you have an executable, make sure to copy it somewhere on your path like `/usr/local/bin` or `C:/Program Files/`.
If on a \*nix/mac system, make sure to run `chmod +x /path/to/crank`.

### Homebrew
This app is also avalable from this [homebrew tap](https://github.com/gesquive/homebrew-tap). Just install the tap and then the app will be available.
```shell
$ brew tap gesquive/tap
$ brew install crank
```

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

## Shell Completion
If you install via homebrew, the shell completion scripts are installed for you. Otherwise, this application can generate its own shell completion scripts for bash/fish/zsh.

### Bash
To generate and load completion scripts for bash
```bash
source <(crank completion bash)
```

To load completions for each session, execute once:

Linux:
```bash
crank completion bash > /etc/bash_completion.d/crank
```
macOS:
```bash
crank completion bash > $(brew --prefix)/etc/bash_completion.d/crank
```

### zsh
If shell completion is not already enabled in your environment, you will need to enable it. You can execute the following once:

```zsh
echo "autoload -U compinit; compinit" >> ~/.zshrc
```

To load completions for each session, execute once:
```zsh
crank completion zsh > "${fpath[1]}/_crank"
```

You will need to start a new shell for this setup to take effect.

### fish
```shell
crank completion fish | source
```

To load completions for each session, execute once:
```shell
crank completion fish > ~/.config/fish/completions/crank.fish
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
