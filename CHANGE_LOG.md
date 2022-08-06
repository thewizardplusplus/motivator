# Change Log

## [v1.2.2](https://github.com/thewizardplusplus/motivator/tree/v1.2.2) (2022-08-06)

Perform refactoring, cover the code with unit tests, improve the output of the `status` command.

- fix the format:
  - of generated task names;
  - of duplicated task names;
  - of the output of the `status` command:
    - improve status displaying;
    - colorize the output;
- correct the text according to the English grammar;
- misc.:
  - perform refactoring;
  - cover the code with unit tests;
  - set up the linters;
  - set up the Travis CI.

## [v1.2.1](https://github.com/thewizardplusplus/motivator/tree/v1.2.1) (2022-07-24)

Use a delay relative to the last displaying for displaying notifications on a schedule; make optional automatic generation of a task name; support for hiding the app name.

- use for displaying notifications on a schedule:
  - a delay relative to the last displaying:
    - use the format of the [`time.ParseDuration()`](https://pkg.go.dev/time@go1.18#ParseDuration) function;
- make optional:
  - automatic generation of a task name, if it was not specified;
  - add a sequential number to duplicated task names;
  - support for seconds in the [cron](https://en.wikipedia.org/wiki/Cron) specification;
- support for hiding the app name.

## [v1.2.0](https://github.com/thewizardplusplus/motivator/tree/v1.2.0) (2022-04-05)

Support for several different tasks for displaying notifications; support for displaying an icon for each notification; support for variable substitution in notifications.

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
- support for variable substitution in notifications:
  - use the format of the [`os.Expand()`](https://pkg.go.dev/os@go1.18#Expand) function;
  - use one common list of variables to substitute in all notifications of all tasks;
- select the app icon (![](docs/logo/logo.png)).

## [v1.1.0](https://github.com/thewizardplusplus/motivator/tree/v1.1.0) (2022-04-05)

Built-in support for running in the background.

- built-in support for running in the background:
  - the console command for starting and restarting;
  - the console command for stopping;
  - the console command for checking of the current running status.

## [v1.0.0](https://github.com/thewizardplusplus/motivator/tree/v1.0.0) (2022-03-26)

Major version.
