package sample

import "github.com/Imomali1/gRPC/pcbook/pb"

// NewKeyboard returns a new keyboard sample
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  RandKeyboardLayout(),
		Backlit: RandBool(),
	}
	return keyboard
}

// NewCPU returns a new CPU sample
func NewCPU() *pb.CPU {
	brand := RandCPUBrand()
	name := RandCPUName(brand)
	numCores := RandInt(2, 8)
	numThreads := RandInt(numCores, 12)
	minFreq := RandFloat(2.0, 5.0)
	maxFreq := RandFloat(minFreq, 8.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numCores),
		NumberThreads: uint32(numThreads),
		MinFreq:       minFreq,
		MaxFreq:       maxFreq,
	}
	return cpu
}
