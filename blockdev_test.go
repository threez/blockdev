package blockdev_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/threez/blockdev"
)

func TestFakeList(t *testing.T) {
	assert := assert.New(t)

	// configure to use the test executable (for persitant output)
	blockdev.Lsblk = "./test/lsblk-fake"
	devices, err := blockdev.List(context.Background())
	if err != nil {
		panic(err)
	}

	assert.Len(devices, 28)
	assert.Equal(128, devices[0].Ra)
	assert.Equal("sda1", devices[8].Children[0].Name)
}

// TODO: test rpi

// TODO: test error
