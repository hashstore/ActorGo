package base_test

import (
	"errors"
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
	sets := []*util.StringSet{
		util.NewStringSet("a", "b", "c"),   //abc
		util.NewStringSet("a", "b", "d"),   //abd
		util.NewStringSet("a", "c", "d&d"), //acdd
	}

	tests := []struct {
		name        string
		wantMatches []bool
		wantError   error
	}{
		{
			"a b c",
			//     abc   abd    acdd
			[]bool{true, false, false},
			nil,
		},
		{
			"a &b c",
			//     abc   abd    acdd
			[]bool{true, false, false},
			nil,
		},
		{
			"a (b | c)",
			//     abc   abd    acdd
			[]bool{true, true, true},
			nil,
		},
		{
			"a (d | (!\"d&d\" & !b))",
			//     abc   abd    acdd
			[]bool{false, true, false},
			nil,
		},
		{
			"a b | c",
			nil,
			errors.New("Cannot mix AND:& and OR:| operations in same expression"),
		},
		{
			"a | b c",
			[]bool{true, true, true},
			nil,
		},
		{
			"a | b & c",
			nil,
			errors.New("Cannot mix AND:& and OR:| operations in same expression"),
		},
		{
			"x (a | b & c)",
			nil,
			errors.New("Cannot mix AND:& and OR:| operations in same expression"),
		},
		{
			"a | (b & c",
			nil,
			errors.New("too few closing parentesises"),
		},
		{
			"a | (b & c))",
			nil,
			errors.New("too many closing parentesises at token:7"),
		},
		{
			"a | (b & c))((a b)",
			nil,
			errors.New("too many closing parentesises at token:7"),
		}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tm, err := base.ParseTagMatch(tt.name)
			if err != nil {
				require.Equal(t, tt.wantError.Error(), err.Error())
			} else {
				require.Equal(t, tt.wantError, nil)
				for i, want := range tt.wantMatches {
					if got := tm.MatchTagSet(sets[i]); got != want {
						t.Errorf("%d:TagMatch.MatchTagSet() = %v, want %v", i, got, want)
					}
				}
			}
		})
	}
}
