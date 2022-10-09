package sample

import (
	"github.com/Imomali1/gRPC/pcbook/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewKeyboard returns a new keyboard sample
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randKeyboardLayout(),
		Backlit: randBool(),
	}

	return keyboard
}

// NewCPU returns a new CPU sample
func NewCPU() *pb.CPU {
	brand := randCPUBrand()
	name := randCPUName(brand)
	numCores := randInt(2, 8)
	numThreads := randInt(numCores, 12)
	minFreq := randFloat64(2.0, 5.0)
	maxFreq := randFloat64(minFreq, 8.0)

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

// NewGPU returns a new GPU sample
func NewGPU() *pb.GPU {
	brand := randGPUBrand()
	name := randGPUName(brand)
	minFreq := randFloat64(1.0, 1.5)
	maxFreq := randFloat64(minFreq, 2.0)

	memory := &pb.Memory{
		Value: uint64(randInt(2, 6)),
		Unit:  pb.Memory_GIGABYTE,
	}

	gpu := &pb.GPU{
		Brand:   brand,
		Name:    name,
		MinFreq: minFreq,
		MaxFreq: maxFreq,
		Memory:  memory,
	}

	return gpu
}

// NewRAM returns a new RAM sample
func NewRAM() *pb.Memory {
	ram := &pb.Memory{
		Value: uint64(randInt(4, 64)),
		Unit:  pb.Memory_GIGABYTE,
	}

	return ram
}

// NewSSD returns a new SSD storage sample
func NewSSD() *pb.Storage {
	ssd := &pb.Storage{
		Driver: pb.Storage_SSD,
		Memory: &pb.Memory{
			Value: uint64(randInt(128, 1024)),
			Unit:  pb.Memory_GIGABYTE,
		},
	}

	return ssd
}

// NewHDD returns a new HDD storage sample
func NewHDD() *pb.Storage {
	hdd := &pb.Storage{
		Driver: pb.Storage_HDD,
		Memory: &pb.Memory{
			Value: uint64(randInt(1, 6)),
			Unit:  pb.Memory_TERABYTE,
		},
	}

	return hdd
}

// NewScreen returns a new screen sample
func NewScreen() *pb.Screen {
	screen := &pb.Screen{
		SizeInch:   randFloat32(13, 17),
		Resolution: randScreenResolution(),
		Panel:      randScreenPanel(),
		MultiTouch: randBool(),
	}

	return screen
}

// NewLaptop returns a new laptop sample
func NewLaptop() *pb.Laptop {
	brand := randLaptopBrand()
	name := randLaptopName(brand)

	laptop := &pb.Laptop{
		Id:       randID(),
		Brand:    brand,
		Name:     name,
		Cpu:      NewCPU(),
		Ram:      NewRAM(),
		Gpus:     []*pb.GPU{NewGPU()},
		Storages: []*pb.Storage{NewSSD(), NewHDD()},
		Screen:   NewScreen(),
		Keyboard: NewKeyboard(),
		Weight: &pb.Laptop_WeightKg{
			WeightKg: randFloat64(1.0, 3.0),
		},
		PriceUsd:    randFloat64(1000, 2500),
		ReleaseYear: uint32(randInt(2010, 2022)),
		UpdatedAt:   timestamppb.Now(),
	}

	return laptop
}
