# File Tree

## Introduction

The File Tree Printer is a Go-based utility that mimics the behavior of the Unix tree command, displaying the directory structure along with the files. It's capable of showing file sizes and can differentiate between files and directories for a clearer structure representation.

## Features

* Recursively print the directory and file structure.
* Option to include file sizes in the output.
* Customizable depth of directory traversal.
* Sort files and directories alphabetically for consistent output.

## Getting started

### Prerequisites

Ensure you have Go installed on your system. You can download it from [the official Go website](https://go.dev/dl/).

### Installing

First, clone the repository to your local machine:

```sh
git clone git@github.com:LikimiaD/File-Tree-Printer.git
cd file-tree-printer
make build
```

Then, to run the program, simply execute:

```sh
./tree <path-to-directory> [-f]
```

The `-f` flag is optional and includes files in the tree along with directories. Without the `-f` flag, only directories are displayed.

### Usage

To use this Makefile, ensure you have `build`, `format`, `test` and `docker` installed on your system. Then, you can execute any of the above targets using the make command. For example:

```sh
make test         # Builds and runs the unit tests
make format       # Formats all source files according to the Google style guide
make build        # Assembles the solution into a single file
make docker       # Builds a Docker container inside which the tests are run
```


## Contributing
Contributions to improve this implementation are welcome. Please follow the style guide for Go code and ensure that any additions come with corresponding unit tests.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.