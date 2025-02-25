<h1 align="center" id="title">Proxet</h1>

<p align="center"><img src="https://socialify.git.ci/rivethorn/proxet/image?font=Jost&amp;issues=1&amp;language=1&amp;name=1&amp;owner=1&amp;pattern=Transparent&amp;theme=Auto" alt="project-image" width="60%"></p>

<p id="description" align="center">CLI tool that sets the desired proxy address system-wide</p>

## Table of Contents

- [Introduction](#introduction)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Introduction

Proxet is a simple command-line tool designed to change Fish shell environment variables so the system uses a custom proxy across the board. It allows you to easily set or reset proxy settings with a single command.

## Features

- Set a custom proxy address
- Reset proxy settings
- Simple and easy-to-use CLI

## Installation

To install Proxet, clone the repository and build the Go application:

```sh
git clone https://github.com/rivethorn/proxet.git
cd proxet
go build -o proxet main.go
```

You can run the app from anywhere by running:

```sh
go build && go install
```

from the project directory.

## Usage

To use Proxet, run the following commands:

- Set a proxy address:

```sh
./proxet -a <proxy_address>
```

- Reset proxy settings:

```sh
./proxet -r
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.
