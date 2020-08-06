package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// BoardSuitsType はドメイン層のスーツタイプに変換する
func BoardSuitsType(suitsType api.BoardSuitsType) (board.SuitsType, error) {
	switch suitsType {
	case api.BOARD_SUITS_TYPE_MONO_TONE:
		return board.SuitsTypeMonoTone, nil
	case api.BOARD_SUITS_TYPE_TWO_TONE:
		return board.SuitsTypeTwoTone, nil
	case api.BOARD_SUITS_TYPE_RAINBOW:
		return board.SuitsTypeRainbow, nil
	}
	return 0, flserr.Errorf("invalid suits type. %s", suitsType)
}
