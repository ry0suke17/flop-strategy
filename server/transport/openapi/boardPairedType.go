package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// BoardPairedType はドメイン層のペアタイプに変換する。
func BoardPairedType(pairedType api.BoardPairedType) (board.PairedType, error) {
	switch pairedType {
	case api.BOARD_PAIRED_TYPE_UNPAIRED:
		return board.PairedTypeUnpaired, nil
	case api.BOARD_PAIRED_TYPE_PAIRED:
		return board.PairedTypePaired, nil
	case api.BOARD_PAIRED_TYPE_TRIPS:
		return board.PairedTypeTrips, nil
	}
	return 0, flserr.Errorf("invalid board paired type. %s", pairedType)
}
