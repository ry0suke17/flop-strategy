package board

// HighCard はボードのハイカードを表す。
type HighCard int

const (
	// HighCardA は A のハイカードを表す。
	HighCardA HighCard = 1
	// HighCardK は K のハイカードを表す。
	HighCardK HighCard = 2
	// HighCardQ は Q のハイカードを表す。
	HighCardQ HighCard = 3
	// HighCardJ は J のハイカードを表す。
	HighCardJ HighCard = 4
	// HighCardT は T のハイカードを表す。
	HighCardT HighCard = 5
	// HighCard8To9 は 8~9 のハイカードを表す。
	HighCard8To9 HighCard = 6
	// HighCard5To7 は 5~7 のハイカードを表す。
	HighCard5To7 HighCard = 7
	// HighCard2To4 は 2~4 のハイカードを表す。
	HighCard2To4 HighCard = 8
)
