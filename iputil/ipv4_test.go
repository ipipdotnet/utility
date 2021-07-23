package iputil

import (
	"net"
	"testing"
)

func TestUInt32ToIPv4(t *testing.T) {
	t.Log(UInt32ToIPv4(1069384235))
}

func TestIPv4ToUInt32(t *testing.T) {
	t.Log(IPv4ToUInt32(net.ParseIP("198.8.8.8")))
}

func TestRangeToCIDR(t *testing.T) {
	RangeToCIDR()
}

func TestIsPrivateIPv4(t *testing.T) {

	t.Log(IsPrivateIPv4(net.IPv4(0, 168, 1, 1)))
	t.Log(IsPrivateIPv4(net.IPv4(127, 168, 1, 1)))
	t.Log(IsPrivateIPv4(net.IPv4(117, 168, 1, 1)))
}

func TestIPv4SubNetworks(t *testing.T) {
	t.Log(IPv4SubNetworks("192.168.0.0/20", 24))
}