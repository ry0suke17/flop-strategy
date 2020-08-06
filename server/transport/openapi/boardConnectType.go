package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// BoardConnectType はドメイン層のコネクトタイプに変換する
func BoardConnectType(connectType api.BoardConnectType) (board.ConnectType, error) {
	switch connectType {
	case api.BOARD_CONNECT_TYPE_DISCONNECT:
		return board.ConnectTypeDisconnected, nil
	case api.BOARD_CONNECT_TYPE_CONNECTED:
		return board.ConnectTypeConnected, nil
	}
	return 0, flserr.Errorf("invalid board connect type. %s", connectType)
}
