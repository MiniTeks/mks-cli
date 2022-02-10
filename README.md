```bash
    ███╗   ███╗██╗  ██╗███████╗     ██████╗██╗     ██╗
    ████╗ ████║██║ ██╔╝██╔════╝    ██╔════╝██║     ██║
    ██╔████╔██║█████╔╝ ███████╗    ██║     ██║     ██║
    ██║╚██╔╝██║██╔═██╗ ╚════██║    ██║     ██║     ██║
    ██║ ╚═╝ ██║██║  ██╗███████║    ╚██████╗███████╗██║
    ╚═╝     ╚═╝╚═╝  ╚═╝╚══════╝     ╚═════╝╚══════╝╚═╝
```

## mks-cli

`mks` is a cli client to interact with [mks-server](https://github.com/MiniTeks/mks-server).

## Build mks-cli from source

### Prerequisites

This application needs a `go compiler` to build and a kubernetes cluster to run.
You can also use `minikube` or `kind` to run on your local machine. Please install
in advance.

### Steps to build

- Clone the repository using:

```bash
    git clone https://github.com/MiniTeks/mks-cli.git
```

- Make sure you have all the go dependencies

```bash
    go mod tidy

    go mod vendor
```

- Build the project using:

```bash
    go build ./cmd/mks
```

- Run as

```bash
    ./mks [options]
```

### Options

```
      --config string   k8s config file (default is ${HOME}/.kube/config) (default "/home/avinkuma/.kube/config")
  -h, --help            help for mks
```

### See Subcommands

- [mks mkspipelinerun](docs/mks_mkspipelinerun.md) - Manage PipelineRuns
- [mks mkstask](docs/mks_mkstask.md) - mkstask <option>
- [mks mkstaskrun](docs/mks_mkstaskrun.md) - Add create list mkstaskrun

### License

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License.
