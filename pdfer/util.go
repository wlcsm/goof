package pdfer

func digitToBuf(d int, buf []byte) int {
	// This is just easier to handle manually. Otherwise we have to deal with
	// the case when s becomes zero and get a division by zero error.
	if d == 0 {
		buf[0] = '0'
		return 1
	}

	// Adjust "s" so we can find the leading digit.
	s := 1000000000
	for 0 == d/s {
		s /= 10
	}

	i := 0
	for ; s > 0; i += 1 {
		leadDigit := d / s
		buf[i] = byte(leadDigit + 48)
		d -= leadDigit * s
		s /= 10
	}

	return i
}
