package add_test

import (
	"testing"

	"github.com/Dynatrace-Reinhard-Pilz/vs-code-sample/add"
)

func TestAdd(t *testing.T) {
	t.Log("Testing add.Add")
	res := add.Add(3, 2)
	t.Log("  ", "result", res)
	if res != 5 {
		t.Fail()
	}
}
