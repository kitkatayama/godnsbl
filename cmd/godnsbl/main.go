package main

import (
	"encoding/binary"
	"encoding/json"
	"net"

	"fmt"
	"log"
	"os"
	"sync"

	"github.com/mrichman/godnsbl"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Please supply a domain name or IP address.")
		os.Exit(1)
	}

	ip, netmask, _ := net.ParseCIDR(os.Args[2])
	bl := os.Args[1]
	ipHead := ip.To4()
	num := int(^binary.BigEndian.Uint32(netmask.Mask) + 1)
	wg := &sync.WaitGroup{}

	results := make([]godnsbl.Result, num)

	for i := 0; i < num; i++ {
		ip := make(net.IP, len(ipHead))
		copy(ip, ipHead)

		binary.BigEndian.PutUint32(ip, binary.BigEndian.Uint32(ip.To4())+uint32(i))
		wg.Add(1)
		go func(i int, ip net.IP) {
			defer wg.Done()
			rbl := godnsbl.Lookup(bl, ip.String())
			if len(rbl.Results) == 0 {
				results[i] = godnsbl.Result{}
			} else {
				results[i] = rbl.Results[0]
			}
		}(i, ip)
	}
	wg.Wait()

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(&results); err != nil {
		log.Println(err)
	}
}
