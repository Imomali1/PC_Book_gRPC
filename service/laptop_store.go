package service

import (
	"errors"
	"fmt"
	"github.com/Imomali1/gRPC/pcbook/pb"
	"github.com/jinzhu/copier"
	"sync"
)

// ErrAlreadyExists is returned when a record with the same ID already exists in memory
var ErrAlreadyExists = errors.New("record already exists")

// LaptopStore is an interface to store laptop
type LaptopStore interface {
	Save(laptop *pb.Laptop) error
}

// InMemoryLaptopStore stores laptop in memory
type InMemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*pb.Laptop
}

// NewLaptopStore returns a new InMemoryLaptopStore
func NewLaptopStore() *InMemoryLaptopStore {
	return &InMemoryLaptopStore{
		data: make(map[string]*pb.Laptop),
	}
}

func (store *InMemoryLaptopStore) Save(laptop *pb.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	tmp := &pb.Laptop{}
	err := copier.Copy(laptop, tmp)
	if err != nil {
		return fmt.Errorf("cannot copy laptop data: %w", err)
	}

	store.data[laptop.Id] = tmp
	return nil
}

/*
To deal with DB
type DBLaptopStore struct {}
*/
