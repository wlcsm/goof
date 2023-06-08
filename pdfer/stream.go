package pdfer

import (
	"fmt"
	"io"
)

type Stream struct {
	Font      string
	FontSize  int
	LocationX int
	LocationY int
	Text      []byte
}

func (s Stream) Dump(w io.Writer, ref int) (n int, err error) {
	body := fmt.Sprintf(`BT
  /%s %d Tf
  %d %d Td
  (%s) Tj
ET`, s.Font, s.FontSize, s.LocationX, s.LocationY, string(s.Text))

	return fmt.Fprintf(w, `%d 0 obj
<< /Length %d >>
stream
%s
endstream
endobj
`, ref, len(body), body)
}
