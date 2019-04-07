# ocp-mig-test-data-cli

CLI for ocp-mig-test-data

## Installation

Add the following line to your ~/.bash_profile:

```export PATH=$GOPATH/bin:$PATH```

Save and exit your editor. Then, source your ~/.bash_profile.

```source ~/.bash_profile```

Download and build CLI

```go get github.com/alexander-demichev/mig-data-cli```

```go install github.com/alexander-demichev/mig-data-cli```

## Usage

CLI is able to create samples, backups and restores separatly or in combination with each other. To list all samples use ``mig-data-cli help``.

Run ``mig-data-cli clone`` to clone repo with sample data ansible playbooks.

```console
$ mig-data-cli route
$ mig-data-cli route --all/-a
$ mig-data-cli route --sample/-s
$ mig-data-cli route --backup/-b
$ mig-data-cli route --restore/-r
$ mig-data-cli route -b -r
```

```console
$ mig-data-cli <sample_name> [FLAGS]
$ mig-data-cli help
$ mig-data-cli list # list all samples
```

## Dev setup

This project requires Go to be installed. On OS X with Homebrew you can just run `brew install go`.

Running it then should be as simple as:

```console
$ make
$ ./bin/mig-data-cli
```

### Testing

``make test``
