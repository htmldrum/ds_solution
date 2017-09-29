package writers

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func NewStdOut(preamble string, postamble string, headings []string) StdOut {
	w := bufio.NewWriter(os.Stdout)
	return StdOut{preamble, postamble, headings, w}
}

type StdOut struct {
	preamble  string
	postamble string
	headings  []string
	fd        *bufio.Writer
}

func (c StdOut) Write(rows []RowPrintable) error {
	formatted_string := c.formatRows(rows)
	r := strings.NewReader(formatted_string)

	if _, err := io.Copy(c.fd, r); err != nil {
		return err
	}

	return c.fd.Flush()
}

func (c StdOut) formatRows(rows []RowPrintable) string {
	f := c.preamble
	f = f + "\n"

	f = f + strings.Join(c.headings, "\t")
	f = f + "\n"

	for _, r := range rows {
		fields := r.Row()
		out_row := strings.Join(fields, "\t")

		f = f + out_row
		f = f + "\n"
	}

	f = f + c.postamble

	return f
}
