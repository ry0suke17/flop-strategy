package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/playerposition"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// PlayerPostion はドメイン層のプレイヤーポジションに変換する
func PlayerPostion(position api.PlayerPosition) (playerposition.Position, error) {
	switch position {
	case api.PLAYER_POSITION_SB:
		return playerposition.PositionSB, nil
	case api.PLAYER_POSITION_BB:
		return playerposition.PositionBB, nil
	case api.PLAYER_POSITION_UTG:
		return playerposition.PositionUTG, nil
	case api.PLAYER_POSITION_HJ:
		return playerposition.PositionHJ, nil
	case api.PLAYER_POSITION_CO:
		return playerposition.PositionCO, nil
	case api.PLAYER_POSITION_BTN:
		return playerposition.PositionBTN, nil
	}
	return "", flserr.Errorf("invalid player position. %s", position)
}
