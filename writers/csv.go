package writers

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func NewCSV(preamble string, postamble string, headings []string, path string) (error, CSV) {
	f, err := os.OpenFile(path, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		return err, CSV{}
	}
	w := bufio.NewWriter(f)

	return nil, CSV{preamble, postamble, headings, w}
}

type CSV struct {
	preamble  string
	postamble string
	headings  []string
	fd        *bufio.Writer
}

func (c CSV) Write(rows []RowPrintable) error {
	formatted_string := c.formatRows(rows)
	r := strings.NewReader(formatted_string)

	if _, err := io.Copy(c.fd, r); err != nil {
		return err
	}

	return c.fd.Flush()
}

func (c CSV) formatRows(rows []RowPrintable) string {
	f := c.preamble
	f = f + "\n"

	f = f + strings.Join(c.headings, ",")
	f = f + "\n"

	for _, r := range rows {
		fields := r.Row()
		out_row := strings.Join(fields, ",")
		f = f + out_row
		f = f + "\n"
	}

	f = f + c.postamble
	return f
}
