package serializer_test

import (
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/hashstore/hashlogic/base"
	"github.com/hashstore/hashlogic/sample"
	"github.com/hashstore/hashlogic/serializer"
	"github.com/stretchr/testify/require"
)

func TestFileSerializer(t *testing.T) {
	t.Parallel()

	binaryFile := "../tmp/matches.bin"
	jsonFile := "../tmp/matches.json"

	match1 := sample.NewTagMatch(3)

	err := serializer.WriteProtobufToBinaryFile(match1, binaryFile)
	require.NoError(t, err)

	err = serializer.WriteProtobufToJSONFile(match1, jsonFile)
	require.NoError(t, err)

	match2 := &base.TagMatch{}
	err = serializer.ReadProtobufFromBinaryFile(binaryFile, match2)
	require.NoError(t, err)

	require.True(t, proto.Equal(match1, match2))
}
