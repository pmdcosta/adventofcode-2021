package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/pmdcosta/adventofcode-2021/pkg/input"
)

func main() {
	lines, err := input.Load("16/input.csv")
	if err != nil {
		log.Fatalf("failed to open input file: %s", err)
	}

	msg := getBITS(lines[0])
	p, _ := getPacket(msg)
	fmt.Println(getVersions(p))
}

func getVersions(pg Packet) int64 {
	switch p := pg.(type) {
	case ValuePacket:
		return p.Version
	case OperatorPacket:
		var sum int64
		for _, ip := range p.Packets {
			sum += getVersions(ip)
		}
		return sum + p.Version
	}
	return 0
}

func getPacket(bits string) (Packet, int) {
	switch getInt(bits[3:6]) {
	case 4:
		p, end := getValuePacket(bits)
		return p, end
	default:
		p, end := getOperatorPacket(bits)
		return p, end
	}
}

func getValuePacket(bits string) (ValuePacket, int) {
	version := getInt(bits[0:3])
	typeID := getInt(bits[3:6])

	var value string
	var last bool
	var i = 6
	for ; !last; i += 5 {
		if getInt(bits[i:i+1]) == 0 {
			last = true
		}
		value = value + bits[i+1:i+5]
	}
	return ValuePacket{Version: version, Type: typeID, Value: getInt(value)}, i
}

func getOperatorPacket(bits string) (OperatorPacket, int) {
	version := getInt(bits[0:3])
	typeID := getInt(bits[3:6])

	var endIndex int
	var packets []Packet
	switch getInt(bits[6:7]) {
	case 0:
		var i = 22
		end := getInt(bits[7:22]) + 22
		for int64(i) < end {
			pk, e := getPacket(bits[i:])
			packets = append(packets, pk)
			i += e
		}
		endIndex = i
	case 1:
		var i = 18
		nPackets := getInt(bits[7:18])
		for p := 0; int64(p) < nPackets; p++ {
			pk, e := getPacket(bits[i:])
			packets = append(packets, pk)
			i += e
		}
		endIndex = i
	}
	return OperatorPacket{Version: version, Type: typeID, Packets: packets}, endIndex
}

func getBITS(hex string) string {
	var bits string
	for _, h := range strings.Split(hex, "") {
		if h == "" || h == "\n" {
			continue
		}
		i, err := strconv.ParseUint(h, 16, 32)
		if err != nil {
			log.Fatal("failed to parse hex:", err)
		}
		bits = bits + fmt.Sprintf("%04b", i)
	}
	return bits
}

func getInt(bits string) int64 {
	i, err := strconv.ParseInt(bits, 2, 64)
	if err != nil {
		log.Fatal("failed to decode bits", err)
	}
	return i
}

type Packet interface {
	GetVersion() int64
}

type ValuePacket struct {
	Version int64
	Type    int64
	Value   int64
}

func (p ValuePacket) GetVersion() int64 {
	return p.Version
}

type OperatorPacket struct {
	Version int64
	Type    int64
	Packets []Packet
}

func (p OperatorPacket) GetVersion() int64 {
	return p.Version
}
