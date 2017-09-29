package writers

import (
	"testing"
)

type RPDouble struct {
	Content []string
}

func (d RPDouble) Row() []string {
	return d.Content
}

func TestFormatRowsAsCSV(t *testing.T) {
	tt := []struct {
		CSV
		Expected     string
		RowPrintable []RowPrintable
	}{
		{
			CSV{
				preamble:  "Cousin Billy",
				postamble: "Give it up for Billy",
				headings:  []string{"Lookit", "Him", "Go"},
			},
			"Cousin Billy\nLookit,Him,Go\nYeahhhhh,I,🤡\nGive it up for Billy",
			[]RowPrintable{
				RPDouble{
					Content: []string{"Yeahhhhh", "I", "🤡"},
				},
			},
		},
	}

	for n, tc := range tt {
		// var b bytes.Buffer
		// tc.CSV.fd = bufio.NewWriter(b)
		// tc.CSV.Write(tt.RPDboule)

		if actual := tc.CSV.formatRows(tc.RowPrintable); tc.Expected != actual {
			t.Errorf("Expected\n%s\n Got\n %s for #%d", tc.Expected, actual, n)
		}
	}
}
