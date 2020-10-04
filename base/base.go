package base

import (
	"fmt"
	"strings"
	"text/scanner"

	"github.com/hashstore/GoActorGo/pb"
)

func checkBlock(tokens []string, start int) []string {
	if tokens[start] != "(" {
		return nil
	}
	level := 1
	for i := start + 1; i < len(tokens); i++ {
		switch tokens[i] {
		case "(":
			level++
		case ")":
			level--
		}
		if level == 0 {
			return tokens[start+1 : i]
		}
	}
	return tokens[start+1:]
}
func parseTokens(tokens []string, negateIt bool) (*pb.TagMatch, error) {
	var matches []*pb.TagMatch
	combineOpDefined := false
	combineWithOr := false
	nextTagNegated := false

	for i := 0; i < len(tokens); i++ {
		block := checkBlock(tokens, i)
		if block != nil {
			match, err := parseTokens(block, nextTagNegated)
			if err != nil {
				return nil, err
			}
			nextTagNegated = false
			matches = append(matches, match)
			i += len(block) + 1
		} else {
			switch tokens[i] {
			case "&":
				if !combineOpDefined {
					combineOpDefined = true
					combineWithOr = false
				} else if combineWithOr {
					return nil, fmt.Errorf("Cannot mix AND:& and OR:| operations in same expression")
				}
			case "|":
				if !combineOpDefined {
					combineOpDefined = true
					combineWithOr = true
				} else if !combineWithOr {
					return nil, fmt.Errorf("Cannot mix AND:& and OR:| operations in same expression")
				}
			case "!":
				nextTagNegated = true
			default:
				if len(matches) > 0 && !combineOpDefined {
					combineOpDefined = true
				}
				matches = append(matches, &pb.TagMatch{
					Tag:    tokens[i],
					Negate: nextTagNegated,
				})
				nextTagNegated = false
			}
		}
	}
	return &pb.TagMatch{
		Matches:     matches,
		CombineAsOr: combineWithOr,
		Negate:      negateIt,
	}, nil
}

// ParseTagMatch parse tag match query
func ParseTagMatch(query string) (*pb.TagMatch, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(query))
	s.Mode = scanner.SkipComments ^ scanner.GoTokens
	s.Error = func(scanner *scanner.Scanner, msg string) {}
	var tokens []string
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}
	// check balance of parentesises
	level := 0
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "(":
			level++
		case ")":
			level--
		}
		if level < 0 {
			return nil, fmt.Errorf("too many closing parentesises at token:%d", i)
		}
	}
	if level != 0 {
		return nil, fmt.Errorf("too few closing parentesises")
	}

	return parseTokens(tokens, false)
}

// func (*pb.TagMatch) testTags(tags ... string) bool {

// }
