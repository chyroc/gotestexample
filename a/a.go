package a

import "github.com/chyroc/gotestexample/b"

func A() int {
	return b.B() + 1
}
