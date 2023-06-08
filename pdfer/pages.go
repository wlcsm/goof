package pdfer

import (
	"fmt"
	"io"
)

type Pages struct {
	Kids     Array
	MediaBox Rectangle
}

func (p Pages) Dump(w io.Writer, ref int) (int, error) {
	return fmt.Fprintf(w, `%d 0 obj
<< /Type /Pages
/Kids %s
/Count %d
/MediaBox %s
>>
endobj
`, ref, p.Kids, len(p.Kids), p.MediaBox)
}
