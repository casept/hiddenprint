### hiddenprint
This is nothing more than a simple utility which outputs a file to `stdout` with its control characters shown in `Printf` format.

### Installation
```shell
go get -v github.com/casept/hiddenprint
```

### Usage
```shell
hiddenprint FILE
```

### Example output
```shell
hiddenprint README.md
### hiddenprint\nThis is nothing more than a simple utility which outputs a file to `stdout` with its control characters shown in `Printf` format.\n\n### Installation\n```shell\ngo get -v github.com/casept/hiddenprint\n```\n\n### Usage\n```shell\nhiddenprint FILE\n```\n
```
