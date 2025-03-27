package day03

import (
	"bytes"
	"fmt"
	"io"
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
	cache  bytes.Buffer
}

func NewDoReader(reader io.Reader) *DoReader {
	return &DoReader{reader: reader, active: true, state: DO_STATE_NONE, cache: bytes.Buffer{}}
}

func (a *DoReader) Read(p []byte) (int, error) {
	n, err := a.reader.Read(p)
	if err != nil {
		return n, err
	}
	a.cache.Reset()
	for i := 0; i < n; i++ {
		if a.state == DO_STATE_NONE {
			if p[i] == 'd' {
				a.state = DO_STATE_D
			} else if a.active {
				a.cache.WriteByte(p[i])
			}
		} else if a.state == DO_STATE_D {
			if p[i] == 'o' {
				a.state = DO_STATE_DO
			} else {
				a.state = DO_STATE_NONE
				a.cache.WriteString("d")
				a.cache.WriteByte(p[i])
			}
		} else if a.state == DO_STATE_DO {
			if p[i] == '(' {
				a.state = DO_STATE_DO_OPEN
			} else if p[i] == 'n' {
				a.state = DO_STATE_DON
			} else {
				a.state = DO_STATE_NONE
				a.cache.WriteString("do")
				a.cache.WriteByte(p[i])
			}
		} else if a.state == DO_STATE_DO_OPEN {
			if p[i] == ')' {
				a.state = DO_STATE_NONE
				a.active = true
			} else {
				a.state = DO_STATE_NONE
				a.cache.WriteString("do(")
				a.cache.WriteByte(p[i])
			}
		} else if a.state == DO_STATE_DON {
			if p[i] == '\'' {
				a.state = DO_STATE_DON_
				a.active = true
			} else {
				a.state = DO_STATE_NONE
				a.cache.WriteString("don")
				a.cache.WriteByte(p[i])
			}
		} else if a.state == DO_STATE_DON_ {
			if p[i] == 't' {
				a.state = DO_STATE_DON_T
				a.active = true
			} else {
				a.state = DO_STATE_NONE
				a.cache.WriteString("don'")
				a.cache.WriteByte(p[i])
			}
		} else if a.state == DO_STATE_DON_T {
			if p[i] == '(' {
				a.state = DO_STATE_DON_T_OPEN
				a.active = true
			} else {
				a.state = DO_STATE_NONE
				a.cache.WriteString("don't")
				a.cache.WriteByte(p[i])
			}
		} else if a.state == DO_STATE_DON_T_OPEN {
			if p[i] == ')' {
				a.state = DO_STATE_NONE
				a.active = false
			} else {
				a.state = DO_STATE_NONE
				a.cache.WriteByte(p[i])
			}
		}
	}

	copy(p, a.cache.Bytes())
	return a.cache.Len(), nil
}

func Part2(in io.Reader) {
	input_data, _ := io.ReadAll(NewMulReader(NewDoReader(in)))
	fmt.Printf("%s\n", input_data)
}
