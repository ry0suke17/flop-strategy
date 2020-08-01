package playerposition

// Position はプレイヤーのポジションを表す。
type Position string

const (
	// PositionSB は SB ポジションを表す
	PositionSB Position = "SB"
	// PositionBB は BB ポジションを表す
	PositionBB Position = "BB"
	// PositionUTG は UTG ポジションを表す
	PositionUTG Position = "UTG"
	// PositionHJ は HJ ポジションを表す
	PositionHJ Position = "HJ"
	// PositionCO は CO ポジションを表す
	PositionCO Position = "CO"
	// PositionBTN は BTN ポジションを表す
	PositionBTN Position = "BTN"
)

func (p Position) String() string { return string(p) }
