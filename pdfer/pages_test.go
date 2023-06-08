package pdfer

import (
	"bytes"
	"testing"
)

func TestDumpPages(t *testing.T) {
	pages := Pages{
		Kids:     []any{reference(3)},
		MediaBox: Rectangle{0, 0, 100, 200},
	}

	b := make([]byte, 0, 100)
	buf := bytes.NewBuffer(b)

	expect := `2 0 obj
<< /Type /Pages
/Kids [3 0 R]
/Count 1
/MediaBox [ 0 0 100 200 ]
>>
endobj
`

	n, err := pages.Dump(buf, 2)
	require(t, err, nil)
	require(t, buf.String(), expect)
	require(t, n, len(expect))
}
