package service1

import "math"
import "encoding/binary"

func (s *Service1) Size() int {
	var size int
	return size
}

func (s *Service1) Marshal(b []byte) int {
	var n int
	return n
}

func (s *Service1) Unmarshal(b []byte) int {
	var n int
	return n
}

func (s *AddIn) Size() int {
	var size int
	size += service1_VarintSize(int64(s.A))
	size += service1_VarintSize(int64(s.B))
	return size
}

func (s *AddIn) Marshal(b []byte) int {
	var n int
	n += binary.PutVarint(b[n:], int64(s.A))
	n += binary.PutVarint(b[n:], int64(s.B))
	return n
}

func (s *AddIn) Unmarshal(b []byte) int {
	var n int
	{
		v, x := binary.Varint(b[n:])
		s.A = int(v)
		n += x
	}
	{
		v, x := binary.Varint(b[n:])
		s.B = int(v)
		n += x
	}
	return n
}

func (s *AddOut) Size() int {
	var size int
	size += service1_VarintSize(int64(s.C))
	return size
}

func (s *AddOut) Marshal(b []byte) int {
	var n int
	n += binary.PutVarint(b[n:], int64(s.C))
	return n
}

func (s *AddOut) Unmarshal(b []byte) int {
	var n int
	{
		v, x := binary.Varint(b[n:])
		s.C = int(v)
		n += x
	}
	return n
}

func service1_UvarintSize(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}
	return i + 1
}

func service1_VarintSize(x int64) int {
	ux := uint64(x) << 1
	if x < 0 {
		ux = ^ux
	}
	return service1_UvarintSize(ux)
}

func service1_GetFloat32(b []byte) float32 {
	return math.Float32frombits(binary.LittleEndian.Uint32(b))
}

func service1_PutFloat32(b []byte, v float32) {
	binary.LittleEndian.PutUint32(b, math.Float32bits(v))
}

func service1_GetFloat64(b []byte) float64 {
	return math.Float64frombits(binary.LittleEndian.Uint64(b))
}

func service1_PutFloat64(b []byte, v float64) {
	binary.LittleEndian.PutUint64(b, math.Float64bits(v))
}
