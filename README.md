![@observerly:iris](https://user-images.githubusercontent.com/84131395/205911224-6f851bb2-28a6-4e7b-8ae3-97c096c6d3eb.png)

---

![GitHub go.mod Go version (branch & subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/observerly/iris/main?filename=go.mod&label=Go)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/observerly/iris)](https://pkg.go.dev/github.com/observerly/iris)
[![Go Report Card](https://goreportcard.com/badge/github.com/observerly/iris)](https://goreportcard.com/report/github.com/observerly/iris)
[![IRIS Actions Status](https://github.com/observerly/iris/actions/workflows/ci.yml/badge.svg)](https://github.com/observerly/iris/actions/workflows/ci.yml)

Iris is observerly's zero-dependency, pure Go library for interoperating with ASCOM Alpaca exposure data structures and the FITS astronomical image format.

---

## Installation

Make sure you have Go installed. Version '1.17.x', or higher is required, and has been tested against.

Simply initialize your project by creating a folder and then running `go mod init` github.com/your/repo inside the repository. 

Then install Iris with the go get command:

```bash
go get -u github.com/observerly/iris
```