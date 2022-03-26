# Motivator

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/motivator?status.svg)](https://godoc.org/github.com/thewizardplusplus/motivator)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/motivator)](https://goreportcard.com/report/github.com/thewizardplusplus/motivator)

The utility for repeatedly displaying notifications in the standard OS way.

## Features

- use the [cron](https://en.wikipedia.org/wiki/Cron) specification for displaying notifications on a schedule:
  - support for seconds in the [cron](https://en.wikipedia.org/wiki/Cron) specification;
- random selection of a notification for displaying;
- support for the [Spintax](https://postmaker.io/blog/spintax-guide/) format in notifications.

## Installation

```
$ go install github.com/thewizardplusplus/motivator@latest
```

## Usage

```
$ motivator -h | -help | --help
$ motivator [options]
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit;
- `-config PATH` &mdash; the path to a config file (default: `config.json`).

## Config

Format of the config in the JSON Schema format: [docs/config.schema.json](docs/config.schema.json).

Example of the config: [config.json](config.json).

## License

The MIT License (MIT)

Copyright &copy; 2022 thewizardplusplus
