package pdfer

import (
	"fmt"
	"testing"
)

func TestFormatReferenceArray(t *testing.T) {
	type test struct {
		arr    Array
		expect string
	}

	for name, test := range map[string]test{
		"Empty": {
			arr:    nil,
			expect: "[]",
		},
		"One element": {
			arr:    []any{reference(1)},
			expect: "[1 0 R]",
		},
		"Multiple elements": {
			arr:    []any{reference(1), reference(2)},
			expect: "[1 0 R 2 0 R]",
		},
	} {
		t.Run(name, func(t *testing.T) {
			res := fmt.Sprintf("%s", test.arr)
			require(t, res, test.expect)
		})
	}
}
