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
	assert.NoError(err)
	assert.Len(devices, 28)
	assert.Equal(128, devices[0].Ra)
	assert.Equal("sda1", devices[8].Children[0].Name)
}

func TestRpiList(t *testing.T) {
	assert := assert.New(t)

	// configure to use the test executable (for persitant output)
	blockdev.Lsblk = "./test/lsblk-rpi"
	devices, err := blockdev.List(context.Background())
	assert.NoError(err)
	assert.Len(devices, 2)
}

func TestErrList(t *testing.T) {
	assert := assert.New(t)

	// configure to use the test executable (for persitant output)
	blockdev.Lsblk = "./test/lsblk-error"
	devices, err := blockdev.List(context.Background())
	assert.Error(err)
	assert.Nil(devices)
}

func TestInvalidList(t *testing.T) {
	assert := assert.New(t)

	// configure to use the test executable (for persitant output)
	blockdev.Lsblk = "./test/lsblk-invalid"
	devices, err := blockdev.List(context.Background())
	assert.Error(err)
	assert.Nil(devices)
}
