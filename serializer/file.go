package serializer

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"os"
)

// WriteProtoBufToJSON writes protocol buffer message to json file
func WriteProtoBufToJSON(message proto.Message, filename string) error {
	data, err := ProtoBufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot convert proto message to json: %w", err)
	}

	err = os.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write json data to file: %w", err)
	}

	return nil
}

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

// ReadProtoBufFromBin reads from file and writes it into proto message
func ReadProtoBufFromBin(filename string, message proto.Message) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary data from file: %w", err)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary data into proto message: %w", err)
	}

	return nil
}
