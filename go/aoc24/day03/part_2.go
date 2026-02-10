package day03

import (
	"io"
	"strconv"
)

type DoReader struct {
	reader   io.Reader
	active   bool
	matchBuf []byte
	overflow []byte
}

func NewDoReader(reader io.Reader) *DoReader {
	return &DoReader{reader: reader, active: true}
}

func (a *DoReader) checkMatch(b byte) (bool, bool) {
	n := len(a.matchBuf)
	// do() : d o ( )
	// don't() : d o n ' t ( )

	switch n {
	case 0:
		if b == 'd' {
			return true, false
		}
	case 1:
		if b == 'o' {
			return true, false
		}
	case 2:
		if b == '(' || b == 'n' {
			return true, false
		}
	case 3:
		if a.matchBuf[2] == '(' {
			// do(
			if b == ')' {
				return true, true
			}
		} else { // 'n'
			// don
			if b == '\'' {
				return true, false
			}
		}
	case 4:
		// don'
		if b == 't' {
			return true, false
		}
	case 5:
		// don't
		if b == '(' {
			return true, false
		}
	case 6:
		// don't(
		if b == ')' {
			return true, true
		}
	}
	return false, false
}

func (a *DoReader) Read(p []byte) (int, error) {
	if len(a.overflow) > 0 {
		n := copy(p, a.overflow)
		a.overflow = a.overflow[n:]
		return n, nil
	}

	n, err := a.reader.Read(p)
	if n == 0 {
		if len(a.matchBuf) > 0 {
			if a.active {
				a.overflow = append(a.overflow, a.matchBuf...)
			}
			a.matchBuf = nil
			if len(a.overflow) > 0 {
				return a.Read(p)
			}
		}
		return 0, err
	}

	var output []byte
	// If p is large, maybe preallocate. But p might be reused.
	// We'll trust append to be reasonable or use a static buffer if performance needed.

	for _, b := range p[:n] {
		match, complete := a.checkMatch(b)
		if match {
			a.matchBuf = append(a.matchBuf, b)
			if complete {
				if len(a.matchBuf) == 4 { // do()
					a.active = true
				} else { // don't()
					a.active = false
				}
				a.matchBuf = a.matchBuf[:0]
			}
		} else {
			// Mismatch
			if a.active {
				output = append(output, a.matchBuf...)
			}
			a.matchBuf = a.matchBuf[:0]

			// Re-check current byte as start of new sequence
			match, _ := a.checkMatch(b)
			if match {
				a.matchBuf = append(a.matchBuf, b)
			} else if a.active {
				output = append(output, b)
			}
		}
	}

	a.overflow = append(a.overflow, output...)
	copied := copy(p, a.overflow)
	a.overflow = a.overflow[copied:]
	return copied, err
}

func Part2(in io.Reader) (int, error) {
	input_data, err := io.ReadAll(newMulReader(NewDoReader(in)))
	if err != nil {
		return 0, err
	}
	result, err := strconv.Atoi(string(input_data))
	if err != nil {
		return 0, err
	}
	return result, nil
}
