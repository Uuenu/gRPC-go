package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// WriteProtobufToJSONFile write protocol Buffer message to JSON file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to JSON: %w", err)
	}

	err2 := ioutil.WriteFile(filename, []byte(data), 0644)
	if err2 != nil {
		return fmt.Errorf("cannot wrile JSON data to file: %w", err)
	}

	return nil
}

// WriteProtobufToBinaryFile writes protocol Buffer message to binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to binary: %w", err)
	}

	if err := ioutil.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("cannot write binary data to file: %w", err)
	}

	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary file to data: %w", err)
	}

	if err := proto.Unmarshal(data, message); err != nil {
		return fmt.Errorf("cannot unmarshal binary data to proto message")
	}

	return nil
}
