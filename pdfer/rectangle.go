package pdfer

import "fmt"

type Rectangle struct {
	x1, y1, x2, y2 int
}

// This is probably way too optimised but oh well.
func (r Rectangle) Format(f fmt.State, c rune) {
	// it will be enough to hold four int and the square brackets
	buf := make([]byte, 50)

	n := 2
	buf[0] = '['
	buf[1] = ' '
	n += digitToBuf(r.x1, buf[2:])
	buf[n] = ' '
	n += 1
	n += digitToBuf(r.y1, buf[n:])
	buf[n] = ' '
	n += 1
	n += digitToBuf(r.x2, buf[n:])
	buf[n] = ' '
	n += 1
	n += digitToBuf(r.y2, buf[n:])
	buf[n] = ' '
	buf[n+1] = ']'
	n += 2

	f.Write(buf[:n])
}
