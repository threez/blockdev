package blockdev_test

import (
	"context"
	"fmt"

	"github.com/threez/blockdev"
)

func ExampleList() {
	// configure to use the test executable (for persitant output)
	blockdev.Lsblk = "./test/lsblk-fake"
	devices, err := blockdev.List(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Devices", devices)
	// Output: Devices [loop0 (7:0) loop1 (7:1) loop2 (7:2) loop3 (7:3) loop4 (7:4) loop5 (7:5) loop6 (7:6) loop7 (7:7) sda (8:0) sr0 (11:0) sr1 (11:1) sr2 (11:2) nbd0 (43:0) nbd1 (43:1) nbd2 (43:2) nbd3 (43:3) nbd4 (43:4) nbd5 (43:5) nbd6 (43:6) nbd7 (43:7) nbd8 (43:8) nbd9 (43:9) nbd10 (43:10) nbd11 (43:11) nbd12 (43:12) nbd13 (43:13) nbd14 (43:14) nbd15 (43:15)]
}
