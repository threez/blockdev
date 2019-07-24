package blockdev

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
)

// Lsblk List information about block devices
var Lsblk = "lsblk"

// BlockDevice contains hierarchical information about linux block devices
type BlockDevice struct {
	// Name device name
	Name string `json:"name"`

	// Kname internal kernel device name
	Kname string `json:"kname"`

	// MajMin major:minor device number
	MajMin string `json:"maj:min"`

	// Fstype filesystem type
	Fstype *string `json:"fstype"`

	// Mountpoint where the device is mounted
	Mountpoint *string `json:"mountpoint"`

	// Label filesystem LABEL
	Label *string `json:"label"`

	// UUID filesystem UUID
	UUID *string `json:"uuid"`

	// Parttype partition type UUID
	Parttype *string `json:"parttype"`

	// Partlabel partition LABEL
	Partlabel *string `json:"partlabel"`

	// Partuuid partition UUID
	Partuuid *string `json:"partuuid"`

	// Partflags partition flags
	Partflags *string `json:"partflags"`

	// Ra read-ahead of the device
	Ra int `json:"ra,string"`

	// Ro read-only device
	Ro int `json:"ro,string"`

	// Rm removable device
	Rm int `json:"rm,string"`

	// Hotplug removable or hotplug device (usb, pcmcia, ...)
	Hotplug int `json:"hotplug,string"`

	// Model device identifier
	Model string `json:"model"`

	// Serial disk serial number
	Serial string `json:"serial"`

	// Size size of the device
	Size int `json:"size,string"`

	// State state of the device
	State *string `json:"state"`

	// Owner user name
	Owner string `json:"owner"`

	// Group group name
	Group string `json:"group"`

	// Mode device node permissions
	Mode string `json:"mode"`

	// Alignment alignment offset
	Alignment int `json:"alignment,string"`

	// MinIo minimum I/O size
	MinIo int `json:"min-io,string"`

	// OptIo optimal I/O size
	OptIo int `json:"opt-io,string"`

	// PhySec physical sector size
	PhySec int `json:"phy-sec,string"`

	// LogSec logical sector size
	LogSec int `json:"log-sec,string"`

	// Rota rotational device
	Rota int `json:"rota,string"`

	// Sched I/O scheduler name
	Sched string `json:"sched"`

	// RqSize request queue size
	RqSize int `json:"rq-size,string"`

	// Type device type
	Type string `json:"type"`

	// DiscAln discard alignment offset
	DiscAln int `json:"disc-aln,string"`

	// DiscGran discard granularity
	DiscGran int `json:"disc-gran,string"`

	// DiscMax discard max bytes
	DiscMax int `json:"disc-max,string"`

	// DiscZero discard zeroes data
	DiscZero int `json:"disc-zero,string"`

	// Wsame write same max bytes
	Wsame int `json:"wsame,string"`

	// Wwn unique storage identifier
	Wwn *string `json:"wwn"`

	// Rand adds randomness
	Rand int `json:"rand,string"`

	// Pkname internal parent kernel device name
	Pkname *string `json:"pkname"`

	// Hctl Host:Channel:Target:Lun for SCSI
	Hctl *string `json:"hctl"`

	// Tran device transport type
	Tran *string `json:"tran"`

	// Subsystems de-duplicated chain of subsystems
	Subsystems string `json:"subsystems"`

	// Rev device revision
	Rev *string `json:"rev"`

	// Vendor device vendor
	Vendor *string `json:"vendor"`

	Children []BlockDevice `json:"children"`
}

func (d BlockDevice) String() string {
	return fmt.Sprintf("%s (%s)", d.Name, d.MajMin)
}

// List of the linux system
func List(ctx context.Context) ([]*BlockDevice, error) {
	// prepare command execution
	cmd := exec.CommandContext(ctx, Lsblk,
		"--all",        // print all devices
		"--json",       // use JSON output format
		"--bytes",      // print SIZE in bytes rather than in human readable format
		"--output-all", // output all columns
	)
	var errorBuffer, resultBuffer bytes.Buffer
	cmd.Stderr = &errorBuffer
	cmd.Stdout = &resultBuffer

	// execute command and return stderr with err in case
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Failed to execute %s: %v; Output: %q", Lsblk, err, errorBuffer.String())
	}

	// parse result
	var result struct {
		BlockDevices []*BlockDevice `json:"blockdevices"`
	}

	err = json.NewDecoder(&resultBuffer).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode the output of %s: %v", Lsblk, err)
	}

	return result.BlockDevices, nil
}
