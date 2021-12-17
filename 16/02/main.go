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
	p, _ := readPacket(msg)
	fmt.Println(evaluatePacket(p))
}

func evaluatePacket(pg Packet) int64 {
	switch p := pg.(type) {
	case ValuePacket:
		return p.Value
	case OperatorPacket:
		switch p.Type {
		case SumType:
			var sum int64
			for _, ip := range p.Packets {
				sum += evaluatePacket(ip)
			}
			return sum
		case ProdType:
			var prod int64 = 1
			for _, ip := range p.Packets {
				prod = prod * evaluatePacket(ip)
			}
			return prod
		case MinType:
			var min int64
			for _, ip := range p.Packets {
				n := evaluatePacket(ip)
				if n < min || min == 0 {
					min = n
				}
			}
			return min
		case MaxType:
			var max int64
			for _, ip := range p.Packets {
				n := evaluatePacket(ip)
				if n > max || max == 0 {
					max = n
				}
			}
			return max
		case LessType:
			var vals []int64
			for _, ip := range p.Packets {
				vals = append(vals, evaluatePacket(ip))
			}
			if vals[0] < vals[1] {
				return 1
			}
		case GreaterType:
			var vals []int64
			for _, ip := range p.Packets {
				vals = append(vals, evaluatePacket(ip))
			}
			if vals[0] > vals[1] {
				return 1
			}
		case EqualType:
			var vals []int64
			for _, ip := range p.Packets {
				vals = append(vals, evaluatePacket(ip))
			}
			if vals[0] == vals[1] {
				return 1
			}
		}
	}
	return 0
}

func readPacket(bits string) (p Packet, end int64) {
	switch Type(getInt(bits[3:6])) {
	case ValueType:
		p, end = readValuePacket(bits)
	case SumType, ProdType, MinType, MaxType, GreaterType, LessType, EqualType:
		p, end = readOperatorPacket(bits)
	default:
		log.Fatal("unknown packet type")
	}
	return p, end
}

func readValuePacket(bits string) (ValuePacket, int64) {
	v, t := readMeta(bits)

	var value string
	var last bool
	var i int64 = 6
	for ; !last; i += 5 {
		if getInt(bits[i:i+1]) == 0 {
			last = true
		}
		value = value + bits[i+1:i+5]
	}
	return ValuePacket{Version: v, Type: t, Value: getInt(value)}, i
}

func readOperatorPacket(bits string) (OperatorPacket, int64) {
	v, t := readMeta(bits)

	var endIndex int64
	var packets []Packet
	switch getInt(bits[6:7]) {
	case 0:
		var i int64 = 22
		for i < getInt(bits[7:22])+22 {
			pk, e := readPacket(bits[i:])
			packets = append(packets, pk)
			i += e
		}
		endIndex = i
	case 1:
		var i int64 = 18
		for p := 0; int64(p) < getInt(bits[7:18]); p++ {
			pk, e := readPacket(bits[i:])
			packets = append(packets, pk)
			i += e
		}
		endIndex = i
	}
	return OperatorPacket{Version: v, Type: t, Packets: packets}, endIndex
}

func readMeta(bits string) (v int64, t Type) {
	return getInt(bits[0:3]), Type(getInt(bits[3:6]))
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

type Packet interface{}

type ValuePacket struct {
	Version int64
	Type    Type

	Value int64
}

type OperatorPacket struct {
	Version int64
	Type    Type

	Packets []Packet
}

type Type int

const (
	SumType     Type = 0
	ProdType    Type = 1
	MinType     Type = 2
	MaxType     Type = 3
	ValueType   Type = 4
	GreaterType Type = 5
	LessType    Type = 6
	EqualType   Type = 7
)
