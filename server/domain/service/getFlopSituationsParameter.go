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
	list, err := s.db.ListFlopSituations(
		ctx,
		heroPosition,
		villainPosition,
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
