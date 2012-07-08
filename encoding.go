package binarydist

// SignMagLittleEndian is the numeric encoding used by the bsdiff tools.
// It implements binary.ByteOrder using a sign-magnitude format
// and little-endian byte order. Only methods Uint64 and String
// have been written; the rest panic.
type signMagLittleEndian struct{}

func (signMagLittleEndian) Uint16(b []byte) uint16 { panic("unimplemented") }

func (signMagLittleEndian) PutUint16(b []byte, v uint16) { panic("unimplemented") }

func (signMagLittleEndian) Uint32(b []byte) uint32 { panic("unimplemented") }

func (signMagLittleEndian) PutUint32(b []byte, v uint32) { panic("unimplemented") }

func (signMagLittleEndian) Uint64(b []byte) uint64 {
	y := int64(b[0]) |
		int64(b[1])<<8 |
		int64(b[2])<<16 |
		int64(b[3])<<24 |
		int64(b[4])<<32 |
		int64(b[5])<<40 |
		int64(b[6])<<48 |
		int64(b[7]&0x7f)<<56

	if b[7]&0x80 != 0 {
		y = -y
	}
	return uint64(y)
}

func (signMagLittleEndian) PutUint64(b []byte, v uint64) { panic("unimplemented") }

func (signMagLittleEndian) String() string { return "signMagLittleEndian" }
