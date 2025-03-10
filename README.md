<h1 align="center" id="title">Proxet</h1>

<p align="center"><img src="https://socialify.git.ci/rivethorn/proxet/image?description=1&font=KoHo&language=1&name=1&owner=1&pattern=Transparent&stargazers=1&theme=Auto" alt="project-image" width="60%"></p>

<p id="description" align="center">CLI tool that sets the desired proxy address system-wide</p>

## Table of Contents

- [Table of Contents](#table-of-contents)
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
- Check to see if the proxy has been set
- Simple and easy-to-use CLI

## Installation

To install Proxet, clone the repository and build the Go application:

```sh
git clone https://github.com/rivethorn/proxet.git
cd proxet
dart compile exe bin/proxet.dart -o proxet
```


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

- Check proxy settings:

```sh
./proxet -c
```

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## License

This project is licensed under the MIT License. See the LICENSE file for details.