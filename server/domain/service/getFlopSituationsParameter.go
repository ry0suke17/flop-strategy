package service

import (
	"context"

	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/domain/model/flopsituationlist"
	"github.com/yneee/flop-strategy/domain/model/playerposition"
	"github.com/yneee/flop-strategy/infra/flserr"
)

// GetFlopSituationsParameter はフロップシチュエーションのパラメータを取得する
func (s *FlopStrtategyService) GetFlopSituationsParameter(
	ctx context.Context,
	heroPosition playerposition.Position,
	villainPosition playerposition.Position,
	potType board.PotType,
	highCard board.HighCard,
	boardPairedType board.PairedType,
	boardSuitsType board.SuitsType,
) ([]*flopsituationlist.Entity, []*flopsituationlist.Entity, bool, error) {
	heroPosEntity, err := s.db.GetPlayerPostion(ctx, heroPosition)
	if err != nil {
		return nil, nil, false, flserr.Wrap(err)
	}
	villainPosEntity, err := s.db.GetPlayerPostion(ctx, villainPosition)
	if err != nil {
		return nil, nil, false, flserr.Wrap(err)
	}
	headsUp := playerposition.HeadsUp{
		HeroPosEntity:    heroPosEntity,
		VillainPosEntity: villainPosEntity,
		HeroPos:          heroPosition,
		VillainPos:       villainPosition,
	}

	getList := func(boardConnectType board.ConnectType) ([]*flopsituationlist.Entity, error) {
		return s.db.ListFlopSituations(
			ctx,
			headsUp.InPosition(),
			headsUp.OutOfPosition(),
			potType,
			highCard,
			boardPairedType,
			boardSuitsType,
			boardConnectType,
		)
	}

	disconnectedList, err := getList(board.ConnectTypeDisconnected)
	if err != nil {
		return nil, nil, false, flserr.Wrap(err)
	}

	// ボードでペアになっていたりトリップスの時は必然的にコネクトにしないのでその場合はリストを取得しない。 {
	var connectedList []*flopsituationlist.Entity
	if boardPairedType == board.PairedTypeUnpaired {
		list, err := getList(board.ConnectTypeConnected)
		if err != nil {
			return nil, nil, false, flserr.Wrap(err)
		}
		connectedList = list
	}
	// }

	return connectedList, disconnectedList, headsUp.IsInPositionHero(), nil
}
