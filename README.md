# blockdev

[![GoDoc](https://godoc.org/github.com/threez/blockdev?status.svg)](https://godoc.org/github.com/threez/blockdev) [![Coverage 100.0%](https://img.shields.io/badge/coverage-100.0-green.svg)]()

List linux block devices in Go.

## Usage

    package main
    
    import (
        "fmt"
	    "context"
        "github.com/threez/blockdev"
    )

    func main() {
        devices, err := blockdev.List(context.Background())
        if err != nil {
            panic(err)
        }

        fmt.Println("Devices", devices)
    }
