package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hashstore/GoActorGo/pb"
	"github.com/hashstore/GoActorGo/sample"
	"github.com/hashstore/GoActorGo/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/matches.bin"
	jsonFile := "../tmp/matches.json"

	laptop1 := sample.NewLaptop()

	err := serializer.WriteProtobufToBinaryFile(laptop1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(laptop1, jsonFile)
	require.NoError(t, err)

	laptop2 := &pb.Laptop{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, laptop2)
	require.NoError(t, err)

	require.True(t, proto.Equal(laptop1, laptop2))
}
