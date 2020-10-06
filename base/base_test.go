package base_test

import (
	"testing"

	"github.com/hashstore/hashlogic/base"
	"github.com/hashstore/hashlogic/util"
	"github.com/stretchr/testify/require"
)

func TestParseTagMatch(t *testing.T) {
	t.Parallel()

	match, err := base.ParseTagMatch("a ( b | !c)")
	if err != nil {
		t.Errorf("Error: %v", err)
	} else {
		require.Equal(t, util.Tokenize("matches:{matches:{tag:\"a\"} matches:{matches:{matches:{tag:\"b\"} matches:{tag:\"c\" negate:true}} combine_as_or:true}}"), util.Tokenize(match.String()))
	}
}

func TestTagMatch_MatchTagSet(t *testing.T) {

	abc := util.NewStringSet("a", "b", "c")
	// abd := util.NewStringSet("a", "b", "d")
	// acd := util.NewStringSet("a", "d", "c")

	type args struct {
		tmText string
		set    *util.StringSet
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"and1",
			args{"a b c", abc},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, err := base.ParseTagMatch(tt.args.tmText)
			if err != nil {
				t.Errorf("Error: %v", err)
			} else {
				if got := tm.MatchTagSet(tt.args.set); got != tt.want {
					t.Errorf("TagMatch.MatchTagSet() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
