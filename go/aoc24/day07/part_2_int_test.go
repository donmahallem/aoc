package day07

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func Test_opConcat(t *testing.T) {
	if i := opConcat(5, 3); i != 53 {
		t.Errorf(`Expected %d and %d to be %d and not %d`, 5, 3, 53, i)
	}
	if i := opConcat(512, 355); i != 512355 {
		t.Errorf(`Expected %d and %d to be %d and not %d`, 512, 355, 512355, i)
	}
}
