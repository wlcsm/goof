package pdfer

import (
	"fmt"
	"testing"
)

func TestRectangle(t *testing.T) {
	rect := Rectangle{0, 0, 100, 200}

	expect := "[ 0 0 100 200 ]"

	res := fmt.Sprintf("%s", rect)
	require(t, res, expect)
}
