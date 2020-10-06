package util_test

import (
	"reflect"
	"testing"

	"github.com/hashstore/hashlogic/util"
	"github.com/stretchr/testify/require"
)

func TestUniformByteMap(t *testing.T) {
	type args struct {
		keys      []string
		defWeight int8
	}
	tests := []struct {
		name string
		args args
		want util.ByteMap
	}{
		{
			"case1",
			args{[]string{"a", "b"}, 2},
			util.ByteMap{"a": 2, "b": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.UniformByteMap(tt.args.keys, tt.args.defWeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniformByteMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUniformWeightMap(t *testing.T) {
	type args struct {
		keys      []string
		defWeight float64
	}
	tests := []struct {
		name string
		args args
		want util.WeightMap
	}{
		{
			"case1",
			args{[]string{"a", "b"}, 2},
			util.WeightMap{"a": 2, "b": 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.UniformWeightMap(tt.args.keys, tt.args.defWeight); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UniformWeightMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStringSet(t *testing.T) {
	type validation struct {
		contains     []string
		absent       []string
		initial_size int
		final_size   int
	}
	tests := []struct {
		name string
		keys []string
		add  []string
		want validation
	}{
		{
			"duplicates",
			[]string{"a", "b", "a"},
			[]string{"b"},
			validation{
				[]string{"b", "a"},
				[]string{"c"},
				2,
				2,
			},
		},
		{
			"just adds",
			[]string{},
			[]string{"a", "b", "a"},
			validation{
				[]string{"b", "a"},
				[]string{"c"},
				0,
				2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := util.NewStringSet(tt.keys...)
			require.Equal(t, got.Size(), tt.want.initial_size)
			for _, k := range tt.add {
				got.Add(k)
			}
			require.Equal(t, got.Size(), tt.want.final_size)
			for _, k := range tt.want.contains {
				require.Equal(t, true, got.Contains(k))
			}
			for _, k := range tt.want.absent {
				require.Equal(t, false, got.Contains(k))
			}
		})
	}
}
