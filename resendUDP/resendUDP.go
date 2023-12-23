package main

import (
	"errors"
	"fmt"
	"log"
	"net"
	"os"
)

// type PacketUDP struct {
// 	Id     uint32
// 	Type   uint8
// 	Length uint16
// 	Body   string
// }

// /*la fonction est degeu+fausse*/
// func initPacketUDP(p *PacketUDP, packet []byte) {
// 	p.Id = binary.BigEndian.Uint32(packet[0:4])
// 	p.Type = packet[4]
// 	p.Length = binary.BigEndian.Uint16(packet[5:7])
// 	p.Body = string(packet[7 : 7+p.Length])
// }

func main() {
	c2, err := net.ListenPacket("udp", ":37086")
	if err != nil {
		log.Fatal("ListenPacket", err)
	}
	c, err := net.ListenPacket("udp", ":9157")
	if err != nil {
		log.Fatal("ListenPacket", err)
	}

	// c.ReadFrom(buf)

	// s := string(buf)

	// udpaddress, err := net.ResolveUDPAddr("udp", s)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	for {
		// c.SetDeadline(time.Now().Add(time.Duration(4000000000)))
		// c2.SetDeadline(time.Now().Add(time.Duration(4000000000)))

		buf := make([]byte, 2048)
		buf2 := make([]byte, 2048)

		size, addr, err := c.ReadFrom(buf)
		if errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Println(err)
			continue
		}
		size2, addr2, err := c2.ReadFrom(buf2)

		if errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Println(err)
			continue
		}

		_, err = c.WriteTo(buf, addr)
		_, err = c2.WriteTo(buf2, addr2)

		if errors.Is(err, os.ErrDeadlineExceeded) {
			fmt.Println(err)
			continue
		}

		// fmt.Println(buf2)

		fmt.Println("Read/Write SUCCESS")
		fmt.Println(addr, buf[:size])
		fmt.Println(addr2, buf2[:size2])
	}

	// fmt.Println(addr, p)

}
