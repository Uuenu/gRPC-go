package serializer

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"main.go/pb"
	"main.go/sample"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}

	err2 := ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err2)
	require.True(t, proto.Equal(laptop1, laptop2))

	err3 := WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err3)
}
