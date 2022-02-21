package add_test

import (
	"testing"

	add "github.com/Dynatrace-Reinhard-Pilz/vs-code-sample"
)

func TestAdd(t *testing.T) {
	res := add.Add(3, 2)
	if res != 5 {
		t.Fail()
	}
}
