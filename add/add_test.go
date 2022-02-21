package add_test

import (
	"testing"

	"github.com/Dynatrace-Reinhard-Pilz/vs-code-sample/add"
)

func TestAdd(t *testing.T) {
	res := add.Add(3, 2)
	if res != 5 {
		t.Fail()
	}
}

func TestSum(t *testing.T) {
	res := add.Sum(3, 2)
	if res != 5 {
		t.Fail()
	}
}
