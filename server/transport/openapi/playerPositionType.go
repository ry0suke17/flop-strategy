package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/playerposition"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// ToAPIPlayerPosition は API で返す用のポシションタイプに変換する。
func ToAPIPlayerPosition(positionType playerposition.PositionType) (api.PlayerPositionType, error) {
	switch positionType {
	case playerposition.PositionTypeInPosition:
		return api.PLAYER_POSITION_TYPE_IN_POSITION, nil
	case playerposition.PositionTypeOutOfPosition:
		return api.PLAYER_POSITION_TYPE_OUT_OF_POSITION, nil
	}
	return "", flserr.Errorf("invalid player position type. %d", positionType)
}
