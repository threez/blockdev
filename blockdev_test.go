package blockdev

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFakeList(t *testing.T) {
	assert := assert.New(t)

	// configure to use the test executable (for persitant output)
	Lsblk = "./test/lsblk-fake"
	devices, err := List(context.Background())
	if err != nil {
		panic(err)
	}

	assert.Len(devices, 28)
	assert.Equal(128, devices[0].Ra)
	assert.Equal("sda1", devices[8].Children[0].Name)
}

// TODO: test rpi

// TODO: test error
