package service

import (
	"context"

	"github.com/yneee/flop-strategy/domain/model/flopsituationlist"

	"github.com/yneee/flop-strategy/infra/flserr"
)

// GetFlopSituationsParameter はフロップシチュエーションのパラメータを取得する
func (s *FlopStrtategyService) GetFlopSituationsParameter(
	ctx context.Context,
	heroPosition string,
	villainPosition string,
	potType string,
	highCard string,
	boardPairedType string,
	boardSuitsType string,
) ([]*flopsituationlist.Entity, error) {
	list, err := s.db.ListFlopSituations(ctx)
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	return list, nil
}
