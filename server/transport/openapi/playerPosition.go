package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/player"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// PlayerPostion はドメイン層のプレイヤーポジションに変換する
func PlayerPostion(position api.PlayerPosition) (player.Position, error) {
	switch position {
	case api.PLAYER_POSITION_SB:
		return player.PositionSB, nil
	case api.PLAYER_POSITION_BB:
		return player.PositionBB, nil
	case api.PLAYER_POSITION_UTG:
		return player.PositionUTG, nil
	case api.PLAYER_POSITION_HJ:
		return player.PositionHJ, nil
	case api.PLAYER_POSITION_CO:
		return player.PositionCO, nil
	case api.PLAYER_POSITION_BTN:
		return player.PositionBTN, nil
	}
	return 0, flserr.Errorf("invalid player position. %s", position)
}
