# blockdev

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
