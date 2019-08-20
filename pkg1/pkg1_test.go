package pkg1

import "testing"

func TestF1(t *testing.T) {
	if F1() != 3 {
		t.Errorf("ERROR 1")
	}
}
