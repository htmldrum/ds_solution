package writers

import (
	"testing"
)

func TestFormatRowsAsStdOut(t *testing.T) {
	tt := []struct {
		StdOut
		Expected     string
		RowPrintable []RowPrintable
	}{
		{
			StdOut{
				preamble:  "Cousin Billy",
				postamble: "Give it up for Billy",
				headings:  []string{"Lookit", "Him", "Go"},
			},
			"Cousin Billy\nLookit\tHim\tGo\nYeahhhhh\tI\t🤡\nGive it up for Billy",
			[]RowPrintable{
				RPDouble{
					Content: []string{"Yeahhhhh", "I", "🤡"},
				},
			},
		},
	}

	for n, tc := range tt {
		if actual := tc.StdOut.formatRows(tc.RowPrintable); tc.Expected != actual {
			t.Errorf("Expected\n%s\n Got\n %s for #%d", tc.Expected, actual, n)
		}
	}

}
