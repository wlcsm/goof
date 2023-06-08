package pdfer

import (
	"bytes"
	"testing"
)

func TestDumpStream(t *testing.T) {
	stream := Stream{
		Font:      "F1",
		FontSize:  20,
		LocationX: 10,
		LocationY: 20,
		Text:      []byte("Hello World"),
	}

	b := make([]byte, 0, 200)
	buf := bytes.NewBuffer(b)

	expect := `4 0 obj
<< /Length 47 >>
stream
BT
  /F1 20 Tf
  10 20 Td
  (Hello World) Tj
ET
endstream
endobj
`
	n, err := stream.Dump(buf, 4)
	require(t, err, nil)
	require(t, buf.String(), expect)
	require(t, n, len(expect))
}
