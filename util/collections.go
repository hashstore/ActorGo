package util

import (
	"strings"
	"text/scanner"
)

//ByteMap minimal association between string key and numerical value
type ByteMap map[string]int8

//UpdateByteMap run function to update map for list of keys
func UpdateByteMap(result ByteMap, mergeOp func(i int, v int8, existed bool) int8, keys ...string) ByteMap {
	for i, k := range keys {
		v, ok := result[k]
		result[k] = mergeOp(i, v, ok)
	}
	return result
}

//UniformByteMap creates uniform ByteMap
func UniformByteMap(keys []string, defWeight int8) ByteMap {
	return UpdateByteMap(ByteMap{}, func(int, int8, bool) int8 { return defWeight }, keys...)
}

//WeightMap maximal association between string key and numerical value
type WeightMap map[string]float64

//UpdateWeightMap run function to update map for list of keys
func UpdateWeightMap(result WeightMap, mergeOp func(i int, v float64, existed bool) float64, keys ...string) WeightMap {
	for i, k := range keys {
		v, ok := result[k]
		result[k] = mergeOp(i, v, ok)
	}
	return result
}

//UniformWeightMap creates uniform WeightMap
func UniformWeightMap(keys []string, defWeight float64) WeightMap {
	return UpdateWeightMap(WeightMap{}, func(int, float64, bool) float64 { return defWeight }, keys...)
}

//StringSet simple set implementation backed by ByteMap
type StringSet struct {
	store ByteMap
}

//Add adds key to set
func (set *StringSet) Add(k string) {
	set.store[k] = 1
}

//Contains check if if string is peresent
func (set *StringSet) Contains(k string) bool {
	return set.store[k] == 1
}

//Size of set
func (set *StringSet) Size() int {
	return len(set.store)
}

//NewStringSet creates StringSet from keys
func NewStringSet(keys ...string) *StringSet {
	return &StringSet{UniformByteMap(keys, 1)}
}

//Tokenize including comments ignore errors
func Tokenize(text string) []string {
	var s scanner.Scanner
	s.Init(strings.NewReader(text))
	s.Mode = scanner.SkipComments ^ scanner.GoTokens
	s.Error = func(s *scanner.Scanner, msg string) {}
	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	return tokens
}
