package iputil

import (
	"encoding/binary"
	"github.com/dspinhirne/netaddr-go"
	"github.com/weborama/cidr"
	"github.com/yl2chen/cidranger"
	"net"
)

var privateIPv4Networks = cidranger.NewPCTrieRanger()

func init() {
	var ipNet *net.IPNet
	_, ipNet, _ = net.ParseCIDR("0.0.0.0/8")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("10.0.0.0/8")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("100.64.0.0/10")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("127.0.0.0/8")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("169.254.0.0/16")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("172.16.0.0/12")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("192.0.0.0/24")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("192.0.2.0/24")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("192.88.99.0/24")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("192.168.0.0/16")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("198.18.0.0/15")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("198.51.100.0/24")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("203.0.113.0/24")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("224.0.0.0/4")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("240.0.0.0/4")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))

	_, ipNet, _ = net.ParseCIDR("255.255.255.255/32")
	privateIPv4Networks.Insert(cidranger.NewBasicRangerEntry(*ipNet))
}

func UInt32ToIPv4(v uint32) net.IP {
	b4 := make([]byte, 4)
	binary.BigEndian.PutUint32(b4, v)

	return net.IP(b4)
}

func IPv4ToUInt32(ipv net.IP) uint32 {
	ip4 := ipv.To4()

	return binary.BigEndian.Uint32(ip4)
}

func CIDRToRange(s string) (startIP, endIP net.IP) {
	return nil, nil
}

func IsPrivateIPv4(ipv net.IP) bool {

	ok, e := privateIPv4Networks.Contains(ipv)
	if e == nil {
		return ok
	} else {
		return false
	}
}

func IPv4SubNetworks(ns string, mask int) ([]string, error) {

	v4n, err := netaddr.ParseIPv4Net(ns)
	if err != nil {
		return nil, err
	}
	sc := v4n.SubnetCount(uint(mask))

	v4s := make([]string, 0, sc)
	for i := 0; i < int(sc); i++ {
		v4c := v4n.NthSubnet(uint(mask), uint32(i))

		v4s = append(v4s, v4c.String())
	}

	return v4s, nil
}

func RangeToCIDR() {// fail
	cidr.IPv4Range2CIDR(
		net.IPv4(255, 255, 255, 0),
		net.IPv4(255, 255, 255, 255),
	)
}