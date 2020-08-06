package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// BoardPairType はドメイン層のペアタイプに変換する。
func BoardPairType(pairType api.BoardPairType) (board.PairType, error) {
	switch pairType {
	case api.BOARD_PAIR_TYPE_UNPAIRED:
		return board.PairTypeUnpaired, nil
	case api.BOARD_PAIR_TYPE_PAIRED:
		return board.PairTypePaired, nil
	case api.BOARD_PAIR_TYPE_TRIPS:
		return board.PairTypeTrips, nil
	}
	return 0, flserr.Errorf("invalid board paired type. %s", pairType)
}
