// Steve Phillips / elimisteve
// 2014.12.13

package help

import (
	"errors"
	"testing"
)

func TestFirstError(t *testing.T) {
	var e1, e2, e3 error = errors.New("e1"), nil, errors.New("e3")
	err := FirstError(e1, e2, e3)
	if err != e1 {
		t.Errorf("Got `%v`, wanted `%v`", err, e1)
	}
}
