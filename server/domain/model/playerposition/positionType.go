package playerposition

// PositionType はポジションのタイプを表す。
type PositionType int

const (
	// PositionTypeInPosition はインポジションを表す。
	PositionTypeInPosition PositionType = 1
	// PositionTypeOutOfPosition はアウトオブポジションを表す。
	PositionTypeOutOfPosition PositionType = 2
)
