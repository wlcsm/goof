package pdfer

import (
	"fmt"
	"io"
)

type PDF struct {
	xref []int
	// The Pages object is written by this library, rather than the
	// user, so we also need to record the reference number of each Page to include
	// it in the Pages object's "Kids" field.
	pages Pages
	n     int
	w     io.Writer
}

func NewPDF(w io.Writer, length, height int) (*PDF, error) {
	n, err := w.Write([]byte("%PDF-1.7\n%¥±ë\n"))
	if err != nil {
		return nil, err
	}

	return &PDF{
		// The first element is reserved for the Catalog, the second is
		// for the Pages which will be written on flush
		xref: []int{n, 0},
		w:    w,
		n:    n,
		pages: Pages{
			MediaBox: Rectangle{0, 0, length, height},
		},
	}, nil
}

type Dumper interface {
	Dump(w io.Writer, ref int) (n int, err error)
}

func (p *PDF) pagesReference() reference {
	return reference(2)
}

// Little messy, but we need to add a method specifically for pages since they
// need to support back references, specifically, the Parent field needs to
// reference the Pages object.
//
// Additionally, the Pages object is written by this library, rather than the
// user, so we also need to record the reference number of each Page to include
// it in the Pages object's "Kids" field.
func (p *PDF) AddPage(page Page) (reference, error) {
	// We just use the length of the xref to obtain the next reference
	// number. Why not, its free
	ref := len(p.xref) + 1

	// TODO we probably shouldn't even allow the user to set the Parent field
	page.parent = p.pagesReference()

	p.xref = append(p.xref, p.n)
	n, err := page.Dump(p.w, ref)
	p.n += n

	p.pages.Kids = append(p.pages.Kids, reference(ref))

	return reference(ref), err
}

func (p *PDF) Add(d Dumper) (reference, error) {
	// We just use the length of the xref to obtain the next reference
	// number. Why not, its free
	ref := len(p.xref) + 1

	p.xref = append(p.xref, p.n)
	n, err := d.Dump(p.w, ref)
	p.n += n

	return reference(ref), err
}

func (p *PDF) Flush() error {
	p.xref[1] = p.n

	n, err := p.pages.Dump(p.w, int(p.pagesReference()))
	if err != nil {
		return err
	}
	p.n += n

	p.xref[0] = p.n
	catalog := Catalog{p.pagesReference()}
	n, err = catalog.Dump(p.w, 1)
	if err != nil {
		return err
	}
	p.n += n

	if err := p.printXref(); err != nil {
		return err
	}

	_, err = fmt.Fprintf(p.w, `trailer
<<  /Root 1 0 R
    /Size %d
>>
startxref
%d
%%%%EOF`, len(p.xref)+1, p.n)
	return err
}

type Catalog struct {
	Pages reference
}

func (c Catalog) Dump(w io.Writer, ref int) (int, error) {
	return fmt.Fprintf(w, `%d 0 obj
<< /Type /Catalog
/Pages %s
>>
endobj
`, ref, c.Pages)
}

func (p *PDF) printXref() error {
	_, err := fmt.Fprintf(p.w, `xref
0 %d
0000000000 65535 f
`, len(p.xref)+1)
	if err != nil {
		return err
	}

	for _, xref := range p.xref {
		_, err := fmt.Fprintf(p.w, "%010d 00000 n\n", xref)
		if err != nil {
			return err
		}
	}
	return nil
}
