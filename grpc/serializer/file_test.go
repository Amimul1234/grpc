package serializer

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/require"
	"grpc/pb/pb"
	"grpc/sample"
	"testing"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/laptop.bin"
	jsonFile := "../tmp/laptop.json"

	laptop1 := sample.NewLaptop()

	err := WriteProtoBufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	err = WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))
}
