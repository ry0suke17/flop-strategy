package board

// PairedType はボードのペアタイプを表す。
type PairedType int

const (
	// PairedTypeUnpaired はペアがないタイプを表す。
	PairedTypeUnpaired PairedType = 1
	// PairedTypePaired はペアであるタイプを表す。
	PairedTypePaired PairedType = 2
	// PairedTypeTrips はトリップすであるタイプを表す。
	PairedTypeTrips PairedType = 3
)
