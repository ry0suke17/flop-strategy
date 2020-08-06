package board

// PairType はボードのペアタイプを表す。
type PairType int

const (
	// PairTypeUnpaired はペアがないタイプを表す。
	PairTypeUnpaired PairType = 1
	// PairTypePaired はペアであるタイプを表す。
	PairTypePaired PairType = 2
	// PairTypeTrips はトリップすであるタイプを表す。
	PairTypeTrips PairType = 3
)
