package a_test

import (
	"github.com/chyroc/gotestexample/a"
	"testing"
)

func TestA(t *testing.T) {
	result := a.A()
	if result != 2 {
		t.Fail()
	}
}
