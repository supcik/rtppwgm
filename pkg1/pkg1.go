package pkg1

import "demo/pkg2"

func F1() int {
	return 1 + pkg2.F2()
}
