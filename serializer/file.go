package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
)

// WriteProtoBufToBin writes protocol buffer message to binary file
func WriteProtoBufToBin(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary:%w", err)
	}
	
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary data to file:%w", err)
	}

	return nil
}
