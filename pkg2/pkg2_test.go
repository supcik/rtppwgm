package pkg2

import "testing"

func TestF2(t *testing.T) {
	if F2() != 2 {
		t.Errorf("ERROR 2")
	}
}
