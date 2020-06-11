package tt

import (
	"testing"
)

func TestT1(t *testing.T) {

	template := `Hello, {{ username | noescape}}!`

	vars := MakeVars()
	vars.Set("username", "wmentor")

	res, err := RenderString(template, vars)
	if err != nil {
		t.Fatal("RenderString failed")
	}

	if string(res) != "Hello, wmentor!" {
		t.Fatal("RenderString return invalid")
	}
}
