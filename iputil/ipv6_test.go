package iputil

import (
	"net"
	"testing"
)

func TestIPv6ToLong(t *testing.T) {

	val := IPv6ToLong(net.ParseIP("2001:250:2001::"))
	t.Log(val)
	t.Log(LongToIPv6(val))
}

func TestIPv6SubNetworks(t *testing.T) {
	t.Log(IPv6SubNetworks("2001:250::/36", 48))
}
