package pdfer

import (
	"fmt"
	"io"
)

type Font struct {
	Name     string
	Subtype  string
	BaseFont string
}

type Resources struct {
	Font Font
}

type Page struct {
	// TODO we probably shouldn't even allow the user to set the Parent field
	parent    reference
	Resources Resources
	Contents  reference
}

func (s Page) Dump(w io.Writer, ref int) (n int, err error) {
	return fmt.Fprintf(w, `%d 0 obj
<< /Type /Page
   /Parent %s
   /Resources
     << /Font
       << /%s
         << /Type /Font
            /Subtype /%s
            /BaseFont /%s
         >>
       >>
     >>
   /Contents %s
>>
endobj
`, ref, s.parent, s.Resources.Font.Name, s.Resources.Font.Subtype, s.Resources.Font.BaseFont, s.Contents)
}
