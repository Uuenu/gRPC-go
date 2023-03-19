package sample

import "main.go/pb"

func NewKeyboard() *pb.Keyboard {
	return &pb.Keyboard{
		Layout:               0,
		Backlit:              false,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     []byte{},
		XXX_sizecache:        0,
	}
}
