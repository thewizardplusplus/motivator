# ![](docs/logo/logo.png) Motivator

[![GoDoc](https://godoc.org/github.com/thewizardplusplus/motivator?status.svg)](https://godoc.org/github.com/thewizardplusplus/motivator)
[![Go Report Card](https://goreportcard.com/badge/github.com/thewizardplusplus/motivator)](https://goreportcard.com/report/github.com/thewizardplusplus/motivator)

The utility for repeatedly displaying notifications in the standard OS way.

## Features

- support for several different tasks for displaying notifications:
  - for each task:
    - support for displaying a task name:
      - automatic generation of a task name, if it was not specified;
      - add a sequential number to duplicated task names;
    - support for displaying an icon for each notification:
      - an icon can be specified for:
        - a notification;
        - a task;
        - the entire config;
      - for each notification, the first specified icon is selected in the order above;
    - use the [cron](https://en.wikipedia.org/wiki/Cron) specification for displaying notifications on a schedule:
      - support for seconds in the [cron](https://en.wikipedia.org/wiki/Cron) specification;
    - random selection of a notification for displaying;
    - support for the [Spintax](https://postmaker.io/blog/spintax-guide/) format in notifications;
- support for variable substitution in notifications:
  - use the format of the [`os.Expand()`](https://pkg.go.dev/os@go1.18#Expand) function;
  - use one common list of variables to substitute in all notifications of all tasks;
- built-in support for running in the background:
  - the console command for starting and restarting;
  - the console command for stopping;
  - the console command for checking of the current running status.

## Installation

```
$ go install github.com/thewizardplusplus/motivator@latest
```

## Usage

```
$ motivator -h | --help
$ motivator start [-c PATH | --config PATH]
$ motivator status
$ motivator stop
$ motivator foreground [-c PATH | --config PATH]
```

Commands:

- `start` &mdash; start (or restart) displaying notifications in the background;
- `status` &mdash; check that notifications are being display in the background;
- `stop` &mdash; stop displaying notifications in the background;
- `foreground` &mdash; start displaying notifications in the foreground.

Options:

- `-h`, `--help` &mdash; show the context-sensitive help;
- `-c PATH`, `--config PATH` &mdash; the path to a config file (default: `config.json`).

## Config

Format of the config in the JSON Schema format: [docs/config.schema.json](docs/config.schema.json).

Example of the config: [docs/config.example.json](docs/config.example.json).

## License

The MIT License (MIT)

Copyright &copy; 2022 thewizardplusplus
