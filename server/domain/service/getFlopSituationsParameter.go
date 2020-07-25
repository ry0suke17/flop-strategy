package service

import (
	"context"

	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/infra/flsopenapi"
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
) (interface{}, error) {
	_, err := s.db.ListFlopSituations(ctx)
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	return &flsopenapi.GetFlopSituationsParameterResponse{
		IpEquity: 0.2,
	}, nil

}
