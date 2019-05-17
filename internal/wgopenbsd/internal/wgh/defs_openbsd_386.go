//+build openbsd,386

// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs defs.go

package wgh

const SIOCGIFGMEMB = 0xc024698a

type Ifgroupreq struct {
	Name  [16]int8
	Len   uint32
	Ifgru [16]byte
}

type Ifgreq struct {
	Ifgrqu [16]byte
}

const SIOCGWGSERV = 0xc03c69c8

type WGGetServ struct {
	Name      [16]int8
	Pubkey    [32]uint8
	Port      uint16
	Num_peers uint32
	Peers     *[32]uint8
}