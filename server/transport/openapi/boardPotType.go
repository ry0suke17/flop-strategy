package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// BoardPotType はドメイン層のポットタイプに変換する
func BoardPotType(potType api.PotType) (board.PotType, error) {
	switch potType {
	case api.POT_TYPE_SRP:
		return board.PotTypeSRP, nil
	case api.POT_TYPE_3_BET:
		return board.PotType3Bet, nil
	case api.POT_TYPE_4_BET:
		return board.PotType4Bet, nil
	}
	return 0, flserr.Errorf("invalid pot type. %s", potType)
}
