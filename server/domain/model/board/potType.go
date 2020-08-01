package board

// PotType はポットのタイプを表す。
type PotType string

const (
	// PotTypeSRP はシングルレイズポットを表す。
	PotTypeSRP PotType = "SRP"
	// PotType3Bet は3ベットポットを表す。
	PotType3Bet PotType = "3BET"
	// PotType4Bet は4ベットポットを表す。
	PotType4Bet PotType = "4BET"
)

func (p PotType) String() string { return string(p) }
