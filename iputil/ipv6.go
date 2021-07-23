package iputil

import (
	"github.com/dspinhirne/netaddr-go"
	"math/big"
	"net"
)

func IPv6ToLong(ipv net.IP) *big.Int {

	b16 := ipv.To16()

	return big.NewInt(0).SetBytes(b16)
}

func LongToIPv6(val *big.Int) net.IP {
	return net.IP(val.Bytes())
}

func IPv6SubNetworks(ns string, mask int) ([]string, error) {

	v6n, err := netaddr.ParseIPv6Net(ns)
	if err != nil {
		return nil, err
	}
	sc := v6n.SubnetCount(uint(mask))

	v6s := make([]string, 0, sc)
	for i := 0; i < int(sc); i++ {
		v6c := v6n.NthSubnet(uint(mask), uint64(i))

		v6s = append(v6s, v6c.String())
	}

	return v6s, nil
}

func INet6ATon() {

}

func INet6NToa() {

}