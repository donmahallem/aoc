package day03

import (
	"io"
	"strconv"
)

type ReaderState int

const (
	DO_STATE_D          ReaderState = 0
	DO_STATE_DO         ReaderState = 1
	DO_STATE_DON                    = 2
	DO_STATE_DON_                   = 3
	DO_STATE_DON_T                  = 4
	DO_STATE_DON_T_OPEN             = 5
	DO_STATE_DO_OPEN                = 6
	DO_STATE_NONE                   = 7
)

type DoReader struct {
	reader io.Reader
	active bool
	state  ReaderState
}

func NewDoReader(reader io.Reader) *DoReader {
	return &DoReader{reader: reader, active: true, state: DO_STATE_NONE}
}

func (a *DoReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	writeIdx := 0
	for i := 0; i < n; i++ {
		char := p[i]
		for {
			advanced := false
			switch a.state {
			case DO_STATE_NONE:
				if char == 'd' {
					a.state = DO_STATE_D
					advanced = true
				} else {
					advanced = true
				}
			case DO_STATE_D:
				if char == 'o' {
					a.state = DO_STATE_DO
					advanced = true
				} else {
					a.state = DO_STATE_NONE
				}
			case DO_STATE_DO:
				if char == '(' {
					a.state = DO_STATE_DO_OPEN
					advanced = true
				} else if char == 'n' {
					a.state = DO_STATE_DON
					advanced = true
				} else {
					a.state = DO_STATE_NONE
				}
			case DO_STATE_DO_OPEN:
				if char == ')' {
					a.state = DO_STATE_NONE
					a.active = true // Found: do()
					advanced = true
				} else {
					a.state = DO_STATE_NONE
				}
			case DO_STATE_DON:
				if char == '\'' {
					a.state = DO_STATE_DON_
					advanced = true
				} else {
					a.state = DO_STATE_NONE
				}
			case DO_STATE_DON_:
				if char == 't' {
					a.state = DO_STATE_DON_T
					advanced = true
				} else {
					a.state = DO_STATE_NONE
				}
			case DO_STATE_DON_T:
				if char == '(' {
					a.state = DO_STATE_DON_T_OPEN
					advanced = true
				} else {
					a.state = DO_STATE_NONE
				}
			case DO_STATE_DON_T_OPEN:
				if char == ')' {
					a.state = DO_STATE_NONE
					a.active = false // Found: don't()
					advanced = true
				} else {
					a.state = DO_STATE_NONE
				}
			default:
				a.state = DO_STATE_NONE
			}

			if advanced {
				break
			}
		}

		if a.active {
			p[writeIdx] = char
			writeIdx++
		}
	}

	return writeIdx, err
}

func Part2(in io.Reader) (int, error) {
	input_data, err := io.ReadAll(NewMulReader(NewDoReader(in)))
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(string(input_data))
	if err != nil {
		return 0, err
	}
	return result, nil
}
