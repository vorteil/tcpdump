package tcpdump

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	name   = "lo"
	filter = "tcp"
)

func TestPrepCapturing(t *testing.T) {

	m := NewPacketCaptureManager(512, DefaultPromiscuousMode, DefaultTimeout)
	m.targetDevice = ""

	_, err := m.prepCapturing()
	assert.Error(t, err)

	m.targetDevice = "123"
	m.capturing = true
	_, err = m.prepCapturing()
	assert.Error(t, err)

}

func TestBasicFunction(t *testing.T) {

	m := NewPacketCaptureManager(512, DefaultPromiscuousMode, DefaultTimeout)

	assert.Equal(t, m.snapshotLen, int32(512), "snapshot length not equal")
	assert.Equal(t, m.promiscuousMode, false, "promisuous bool not equal")
	assert.Equal(t, m.timeout, DefaultTimeout, "timeout not equal")

	m.SetDevice(name)
	assert.Equal(t, m.targetDevice, name, "names not equal")

	m.SetFilter(filter)
	assert.Equal(t, m.bpfFilter, filter, "filter not equal")

}

func TestConfigErrors(t *testing.T) {

	m := NewPacketCaptureManager(DefaultSnapshotLen, DefaultPromiscuousMode, DefaultTimeout)
	m.capturing = true

	err := m.SetDevice(name)
	assert.Error(t, err)

	m.SetFilter(filter)
	assert.Error(t, err)

	m.capturing = false

	err = m.SetDevice(name)
	assert.NoError(t, err)

	m.SetFilter(filter)
	assert.NoError(t, err)

	// we are not root or have capabilities to capture
	_, err = m.prepCapturing()
	assert.Error(t, err)

}
