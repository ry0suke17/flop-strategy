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
	boardPairType board.PairType,
	boardSuitsType board.SuitsType,
	boardConnectType board.ConnectType,
) (flopsituationlist.Entities, playerposition.PositionType, error) {
	// ボードでペアになっていたりトリップスの時は必然的にコネクトにしないのでエラーを返す。 {
	if boardConnectType == board.ConnectTypeConnected && boardPairType != board.PairTypeUnpaired {
		return nil, 0, flserr.Errorf(
			"should specified unpaired when connected. boardConnectType=%d, boardPairType=%d",
			boardConnectType,
			boardPairType,
		)
	}
	// }

	heroPosEntity, err := s.db.GetPlayerPostion(ctx, heroPosition)
	if err != nil {
		return nil, 0, flserr.Wrap(err)
	}
	villainPosEntity, err := s.db.GetPlayerPostion(ctx, villainPosition)
	if err != nil {
		return nil, 0, flserr.Wrap(err)
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
		boardPairType,
		boardSuitsType,
		boardConnectType,
	)
	if err != nil {
		return nil, 0, flserr.Wrap(err)
	}

	return list, headsUp.HeroPositionType(), nil
}
