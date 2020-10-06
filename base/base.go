package base

import (
	"fmt"
	"strconv"

	"github.com/hashstore/hashlogic/util"
)

//TagMatchWithTag creates leaf TagMatch from string
func TagMatchWithTag(text string, negate bool) *TagMatch {
	return &TagMatch{
		TagOrMatches: &TagMatch_Tag{Tag: text},
		Negate:       negate,
	}
}

//TagMatchWithMatches creates OR/AND combination of TagMatch'es
func TagMatchWithMatches(matches []*TagMatch, combineWithOr bool, negateIt bool) *TagMatch {
	return &TagMatch{
		TagOrMatches: &TagMatch_Matches{&TagMatch_TagMatches{Matches: matches}},
		CombineAsOr:  combineWithOr,
		Negate:       negateIt,
	}
}

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
func parseTokens(tokens []string, negateIt bool) (*TagMatch, error) {
	var matches []*TagMatch
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
				text := tokens[i]
				quote := text[0]
				last := len(text) - 1
				if quote == '"' && quote == text[last] {
					u, err := strconv.Unquote(text)
					if err == nil {
						text = u
					}
				}
				matches = append(matches, TagMatchWithTag(text, nextTagNegated))
				nextTagNegated = false
			}
		}
	}
	return TagMatchWithMatches(matches, combineWithOr, negateIt), nil
}

// ParseTagMatch parse tag match query
func ParseTagMatch(query string) (*TagMatch, error) {
	tokens := util.Tokenize(query)
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

//MatchTagSet validates if `TagMatch` rule matches tags
func (tm *TagMatch) MatchTagSet(tagSet *util.StringSet) bool {
	result := false
	switch v := tm.TagOrMatches.(type) {
	case *TagMatch_Tag:
		result = tagSet.Contains(v.Tag)
	case *TagMatch_Matches:
		if tm.CombineAsOr {
			result = false
			for _, m := range v.Matches.Matches {
				if m.MatchTagSet(tagSet) {
					result = true
					break
				}
			}
		} else {
			result = true
			for _, m := range v.Matches.Matches {
				if !m.MatchTagSet(tagSet) {
					result = false
					break
				}
			}
		}
	}
	if tm.Negate {
		result = !result
	}
	return result

}
