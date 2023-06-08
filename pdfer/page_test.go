package pdfer

import (
	"bytes"
	"testing"
)

func TestDumpPage(t *testing.T) {
	page := Page{
		parent: reference(3),
		Resources: Resources{
			Font: Font{
				Name:     "F1",
				Subtype:  "Type1",
				BaseFont: "Times-Roman",
			},
		},
		Contents: reference(2),
	}

	b := make([]byte, 0, 200)
	buf := bytes.NewBuffer(b)

	expect := `4 0 obj
<< /Type /Page
   /Parent 3 0 R
   /Resources
     << /Font
       << /F1
         << /Type /Font
            /Subtype /Type1
            /BaseFont /Times-Roman
         >>
       >>
     >>
   /Contents 2 0 R
>>
endobj
`

	n, err := page.Dump(buf, 4)
	require(t, err, nil)
	require(t, buf.String(), expect)
	require(t, n, len(expect))
}
