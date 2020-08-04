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
) ([]*flopsituationlist.Entity, error) {
	heroPosEntity, err := s.db.GetPlayerPostion(ctx, heroPosition)
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	villainPosEntity, err := s.db.GetPlayerPostion(ctx, villainPosition)
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	headsUp := playerposition.HeadsUp{
		HeroPosEntity:    heroPosEntity,
		VillainPosEntity: villainPosEntity,
		HeroPos:          heroPosition,
		VillainPos:       villainPosition,
	}

	list, err := s.db.ListFlopSituations(
		ctx,
		headsUp.InPosition(),
		headsUp.OutOfPosition(),
		potType,
		highCard,
		boardPairedType,
		boardSuitsType,
	)
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	return list, nil
}
