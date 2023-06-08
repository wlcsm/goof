package pdfer

import "fmt"

type reference int

func (r reference) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, "%d 0 R", int(r))
}

// Questioning whether to make a generic one for all Array element types
type Array []any

// This is probably way too optimised but oh well.
func (r Array) Format(f fmt.State, c rune) {
	if len(r) == 0 {
		f.Write([]byte("[]"))
		return
	}

	f.Write([]byte("["))

	fmt.Fprintf(f, "%s", r[0])
	for _, ref := range r[1:] {
		fmt.Fprintf(f, " %s", ref)
	}

	f.Write([]byte("]"))
}
