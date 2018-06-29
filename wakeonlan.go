package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

func macDie(e string) {
	fmt.Fprintln(os.Stderr, "Bad MAC address format :", e)
	os.Exit(1)
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, os.Args[0], "IPADDR:PORT MACADDR")
		os.Exit(1)
	}
	argIP := os.Args[1]
	argMac := os.Args[2]
	txtMac := strings.Split(argMac, "-")
	if len(txtMac) != 6 {
		txtMac = strings.Split(argMac, ":")
	}
	if len(txtMac) != 6 {
		macDie(argMac)
	}
	mac := make([]byte, 6)
	for i, m := range txtMac {
		u, err := strconv.ParseUint(m, 16, 8)
		if err != nil {
			macDie(err.Error())
		}
		mac[i] = byte(u)
	}
	c, err := net.Dial("udp", argIP)
	if err != nil {
		fmt.Fprintln(os.Stderr, `Can't setup UDP "connection":`, err)
		os.Exit(1)
	}
	const pktLen = 17 * 6
	pkt := make([]byte, 17*6)
	i := 0
	for ; i < 6; i++ {
		pkt[i] = 0xff
	}
	for ; i < pktLen; i += 6 {
		copy(pkt[i:i+6], mac)
	}
	if _, err = c.Write(pkt); err != nil {
		fmt.Fprintln(os.Stderr, "Can't send UDP packet:", err)
		os.Exit(1)
	}
	if err = c.Close(); err != nil {
		fmt.Fprintln(os.Stderr, `Can't close UDP "connection":`, err)
		os.Exit(1)
	}
}
