package day03

import (
	"io"
	"strconv"
	"strings"

	"github.com/donmahallem/aoc/go/aoc_utils/bytes"
)

type MulReaderState int

const (
	MUL_STATE_M              MulReaderState = 0
	MUL_STATE_MU             MulReaderState = 1
	MUL_STATE_MUL                           = 2
	MUL_STATE_MUL_OPEN                      = 3
	MUL_STATE_MUL_OPEN_NUM1                 = 4
	MUL_STATE_MUL_OPEN_COMMA                = 5
	MUL_STATE_MUL_OPEN_NUM2                 = 6
	MUL_STATE_NONE                          = 7
)

type MulReader struct {
	reader       io.Reader
	state        MulReaderState
	num1cache    int
	num2cache    int
	currentSum   int
	resultReader io.Reader
}

func NewMulReader(reader io.Reader) *MulReader {
	return &MulReader{reader: reader, state: MUL_STATE_NONE, num1cache: 0, num2cache: 0, currentSum: 0}
}

func (a *MulReader) Read(p []byte) (int, error) {
	if a.resultReader != nil {
		return a.resultReader.Read(p)
	}

	n, err := a.reader.Read(p)
	for i := 0; i < n; i++ {
		char := p[i]
		for {
			advanced := false
			switch a.state {
			case MUL_STATE_NONE:
				if char == 'm' {
					a.state = MUL_STATE_M
					advanced = true
				} else {
					advanced = true
				}
			case MUL_STATE_M:
				if char == 'u' {
					a.state = MUL_STATE_MU
					advanced = true
				} else {
					a.state = MUL_STATE_NONE
				}
			case MUL_STATE_MU:
				if char == 'l' {
					a.state = MUL_STATE_MUL
					advanced = true
				} else {
					a.state = MUL_STATE_NONE
				}
			case MUL_STATE_MUL:
				if char == '(' {
					a.state = MUL_STATE_MUL_OPEN
					advanced = true
				} else {
					a.state = MUL_STATE_NONE
				}
			case MUL_STATE_MUL_OPEN:
				if val, ok := bytes.ParseIntFromByte[int](char); ok {
					a.state = MUL_STATE_MUL_OPEN_NUM1
					a.num1cache = val
					advanced = true
				} else {
					a.state = MUL_STATE_NONE
				}
			case MUL_STATE_MUL_OPEN_NUM1:
				if val, ok := bytes.ParseIntFromByte[int](char); ok {
					// defensive: prevent int overflow if needed, though inputs are typically small
					a.num1cache = a.num1cache*10 + val
					advanced = true
				} else if char == ',' {
					a.state = MUL_STATE_MUL_OPEN_COMMA
					advanced = true
				} else {
					a.state = MUL_STATE_NONE
				}
			case MUL_STATE_MUL_OPEN_COMMA:
				if val, ok := bytes.ParseIntFromByte[int](char); ok {
					a.state = MUL_STATE_MUL_OPEN_NUM2
					a.num2cache = val
					advanced = true
				} else {
					a.state = MUL_STATE_NONE
				}
			case MUL_STATE_MUL_OPEN_NUM2:
				if val, ok := bytes.ParseIntFromByte[int](char); ok {
					a.num2cache = a.num2cache*10 + val
					advanced = true
				} else if char == ')' {
					a.currentSum += a.num1cache * a.num2cache
					a.state = MUL_STATE_NONE
					advanced = true
				} else {
					a.state = MUL_STATE_NONE
				}
			default:
				a.state = MUL_STATE_NONE
			}

			if advanced {
				break
			}
		}
	}

	if err == io.EOF {
		// When input ends, switch to serving the result.
		a.resultReader = strings.NewReader(strconv.Itoa(a.currentSum))
		return a.resultReader.Read(p)
	}

	// We return 0, nil because we consumed the input but produced no output yet.
	// This is acceptable here as we are treating the Reader as a sink until EOF.
	return 0, nil
}
func Part1(in io.Reader) (int, error) {
	input_data, err := io.ReadAll(NewMulReader(in))
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(string(input_data))
	if err != nil {
		return 0, err
	}
	return result, nil
}
