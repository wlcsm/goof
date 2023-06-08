package pdfer

import "testing"

func TestDigitToBuf(t *testing.T) {
	buf := make([]byte, 10)

	n := digitToBuf(10, buf)
	require(t, n, 2)
	require(t, string(buf[:n]), "10")

	n = digitToBuf(0, buf)
	require(t, n, 1)
	require(t, string(buf[:n]), "0")
}
