package pdfer

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func require[T comparable](t *testing.T, got, expect T) {
	t.Helper()
	if expect != got {
		os.WriteFile("expected.txt", []byte(fmt.Sprintf("%v", expect)), os.ModePerm)
		os.WriteFile("got.txt", []byte(fmt.Sprintf("%v", got)), os.ModePerm)
		t.Fatalf(`
-- expected --
%v
-- got --
%v`, expect, got)
	}
}

func TestEmptyPDF(t *testing.T) {
	b := make([]byte, 0, 1000)
	buf := bytes.NewBuffer(b)

	pdf, err := NewPDF(buf, 300, 144)
	require(t, err, nil)
	err = pdf.Flush()
	require(t, err, nil)

	expect := `%PDF-1.7
%¥±ë
2 0 obj
<< /Type /Pages
/Kids []
/Count 0
/MediaBox [ 0 0 300 144 ]
>>
endobj
1 0 obj
<< /Type /Catalog
/Pages 2 0 R
>>
endobj
xref
0 3
0000000000 65535 f
0000000095 00000 n
0000000017 00000 n
trailer
<<  /Root 1 0 R
    /Size 3
>>
startxref
144
%%EOF`

	require(t, buf.String(), expect)
}

func TestBlankCanvasPDF(t *testing.T) {
	b := make([]byte, 0, 1000)
	buf := bytes.NewBuffer(b)

	pdf, err := NewPDF(buf, 300, 144)
	require(t, err, nil)

	stream := Stream{
		Font:      "F1",
		FontSize:  20,
		LocationX: 10,
		LocationY: 20,
		Text:      []byte("Hello world"),
	}
	streamRef, err := pdf.Add(stream)
	require(t, err, nil)

	page := Page{
		Resources: Resources{
			Font: Font{
				Name:     "F1",
				Subtype:  "Type1",
				BaseFont: "Times-Roman",
			},
		},
		Contents: streamRef,
	}
	_, err = pdf.AddPage(page)
	require(t, err, nil)

	require(t, err, nil)
	err = pdf.Flush()
	require(t, err, nil)

	expect := `%PDF-1.7
%¥±ë
3 0 obj
<< /Length 47 >>
stream
BT
  /F1 20 Tf
  10 20 Td
  (Hello world) Tj
ET
endstream
endobj
4 0 obj
<< /Type /Page
   /Parent 2 0 R
   /Resources
     << /Font
       << /F1
         << /Type /Font
            /Subtype /Type1
            /BaseFont /Times-Roman
         >>
       >>
     >>
   /Contents 3 0 R
>>
endobj
2 0 obj
<< /Type /Pages
/Kids [4 0 R]
/Count 1
/MediaBox [ 0 0 300 144 ]
>>
endobj
1 0 obj
<< /Type /Catalog
/Pages 2 0 R
>>
endobj
xref
0 5
0000000000 65535 f
0000000425 00000 n
0000000342 00000 n
0000000017 00000 n
0000000114 00000 n
trailer
<<  /Root 1 0 R
    /Size 5
>>
startxref
474
%%EOF`

	require(t, buf.String(), expect)
}
