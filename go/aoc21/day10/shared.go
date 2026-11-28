package day10

import (
	"github.com/donmahallem/aoc/go/aoc_utils"
)

var scoring = map[byte]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func parseAst(data []byte) (score int, ast []byte, err error) {

	for _, b := range data {
		switch b {
		case '(', '[', '{', '<':
			ast = append(ast, b)
		case ')', ']', '}', '>':
			if len(ast) == 0 {
				return scoring[b], nil, aoc_utils.NewParseErrorf("unexpected closing bracket: %c", b)
			}
			last := ast[len(ast)-1]
			if (b == ')' && last != '(') ||
				(b == ']' && last != '[') ||
				(b == '}' && last != '{') ||
				(b == '>' && last != '<') {
				return scoring[b], nil, aoc_utils.NewParseErrorf("mismatched closing bracket: %c", b)
			}
			ast = ast[:len(ast)-1]
		default:
			return 0, nil, aoc_utils.NewParseErrorf("invalid character: %c", b)
		}
	}
	return 0, ast, nil
}
