package day03

import (
	"fmt"
	"io"
	"strconv"
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
	reader     io.Reader
	state      MulReaderState
	num1cache  int
	num2cache  int
	currentSum int
}

func NewMulReader(reader io.Reader) *MulReader {
	return &MulReader{reader: reader, state: MUL_STATE_NONE, num1cache: 0, num2cache: 0, currentSum: 0}
}

func (a *MulReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		if err == io.EOF {
			sumBuffer := []byte(strconv.Itoa(a.currentSum))
			copy(p, sumBuffer)
			return len(sumBuffer), err
		}
		return n, err
	}
	for i := 0; i < n; i++ {
		if a.state == MUL_STATE_NONE && p[i] == 'm' {
			a.state = MUL_STATE_M
		} else if a.state == MUL_STATE_M {
			if p[i] == 'u' {
				a.state = MUL_STATE_MU
			} else {
				a.state = MUL_STATE_NONE
			}
		} else if a.state == MUL_STATE_MU {
			if p[i] == 'l' {
				a.state = MUL_STATE_MUL
			} else {
				a.state = MUL_STATE_NONE
			}
		} else if a.state == MUL_STATE_MUL {
			if p[i] == '(' {
				a.state = MUL_STATE_MUL_OPEN
			} else {
				a.state = MUL_STATE_NONE
			}
		} else if a.state == MUL_STATE_MUL_OPEN {
			if p[i] >= '0' && p[i] <= '9' {
				a.state = MUL_STATE_MUL_OPEN_NUM1
				a.num1cache = int(p[i] - '0')
			} else {
				a.state = MUL_STATE_NONE
			}
		} else if a.state == MUL_STATE_MUL_OPEN_NUM1 {
			if p[i] >= '0' && p[i] <= '9' {
				a.num1cache = a.num1cache*10 + int(p[i]-'0')
			} else if p[i] == ',' {
				a.state = MUL_STATE_MUL_OPEN_COMMA
			} else {
				a.state = MUL_STATE_NONE
			}
		} else if a.state == MUL_STATE_MUL_OPEN_COMMA {
			if p[i] >= '0' && p[i] <= '9' {
				a.state = MUL_STATE_MUL_OPEN_NUM2
				a.num2cache = int(p[i] - '0')
			} else {
				a.state = MUL_STATE_NONE
			}
		} else if a.state == MUL_STATE_MUL_OPEN_NUM2 {
			if p[i] >= '0' && p[i] <= '9' {
				a.num2cache = a.num2cache*10 + int(p[i]-'0')
			} else if p[i] == ')' {
				a.currentSum += a.num1cache * a.num2cache
				a.state = MUL_STATE_NONE
			} else {
				a.state = MUL_STATE_NONE
			}
		}
	}
	return 0, nil
}
func Part1(in io.Reader) {
	input_data, _ := io.ReadAll(NewMulReader(in))
	fmt.Printf("%s\n", input_data)
}
