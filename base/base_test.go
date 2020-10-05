package base_test

import (
	"testing"

	"github.com/hashstore/hashlogic/sample"
	"github.com/stretchr/testify/require"
)

func TestParseTagMatch(t *testing.T) {
	t.Parallel()

	match := sample.NewTagMatch(3)

	require.Equal(t, "", match.String())
}
