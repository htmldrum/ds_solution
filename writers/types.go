package writers

type ReportWriter interface {
	Write([]RowPrintable) error
}

type RowPrintable interface {
	Row() []string
}
