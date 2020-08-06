package service

import (
	"context"

	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/domain/model/flopsituationlist"
	"github.com/yneee/flop-strategy/domain/model/playerposition"
)

// DBRepository は DB のリポジトリを表す。
type DBRepository interface {
	// ListFlopSituations はフロップシチュエーションのリストを取得する。
	ListFlopSituations(
		ctx context.Context,
		inPosition playerposition.Position,
		outOfPosition playerposition.Position,
		potType board.PotType,
		highCard board.HighCard,
		boardPairType board.PairType,
		boardSuitsType board.SuitsType,
		boardConnectType board.ConnectType,
	) ([]*flopsituationlist.Entity, error)
	// GetPlayerPostion はプレイヤーのポジションを取得する。
	GetPlayerPostion(
		ctx context.Context,
		position playerposition.Position,
	) (*playerposition.Entity, error)
}

// FlopStrtategyService はフロップシチュエーションを扱うサービスを表す。
type FlopStrtategyService struct {
	db DBRepository
}

// NewFlopStrtategyService は新しいサービスを生成する。
func NewFlopStrtategyService(
	db DBRepository,
) *FlopStrtategyService {
	return &FlopStrtategyService{
		db: db,
	}
}
