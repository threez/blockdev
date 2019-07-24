# blockdev

List linux block devices

## Usage

    package main
    
    import (
        "fmt"
        "github.com/threez/blockdev"
    )

    func main() {
        devices, err := blockdev.List(context.Background())
        if err != nil {
            panic(err)
        }  

        fmt.Println("Devices", devices)
    }
