package playerposition

// HeadsUp はヘッズアップ（ヒーロー、ビランの2つの組み合わせ）を表す。
type HeadsUp struct {
	HeroPosEntity    *Entity
	VillainPosEntity *Entity
	HeroPos          Position
	VillainPos       Position
}

// IsInPositionHero はヒーローがインポジションであるかを取得する。
func (h *HeadsUp) IsInPositionHero() bool {
	return h.HeroPosEntity.PostFlopActionOrder > h.VillainPosEntity.PostFlopActionOrder
}

// InPosition はインポジションの値を返す。
func (h *HeadsUp) InPosition() Position {
	if h.IsInPositionHero() {
		return h.HeroPos
	}
	return h.VillainPos
}

// OutOfPosition はアウトオブポジションの値を返す
func (h *HeadsUp) OutOfPosition() Position {
	if h.IsInPositionHero() {
		return h.VillainPos
	}
	return h.HeroPos
}

// HeroPositionType はヒーローのポジションタイプを返す。
func (h *HeadsUp) HeroPositionType() PositionType {
	if h.IsInPositionHero() {
		return PositionTypeInPosition
	}
	return PositionTypeOutOfPosition
}
