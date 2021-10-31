package main

import (
	"fmt"
	"strings"
)

type IPAddr [4]byte

func (i IPAddr) String() string {
	var s []string
	for _, val := range i {
		s = append(s, fmt.Sprintf("%v", val))
	}

	return strings.Join(s, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
