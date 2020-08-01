package openapi

import (
	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/infra/flserr"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// BoardHighCard はドメイン層のハイカードに変換する
func BoardHighCard(highCard api.HighCard) (board.HighCard, error) {
	switch highCard {
	case api.HIGH_CARD_A:
		return board.HighCardA, nil
	case api.HIGH_CARD_K:
		return board.HighCardK, nil
	case api.HIGH_CARD_Q:
		return board.HighCardQ, nil
	case api.HIGH_CARD_J:
		return board.HighCardJ, nil
	case api.HIGH_CARD_T:
		return board.HighCardT, nil
	case api.HIGH_CARD_8_TO_9:
		return board.HighCard8To9, nil
	case api.HIGH_CARD_5_TO_7:
		return board.HighCard5To7, nil
	case api.HIGH_CARD_2_TO_4:
		return board.HighCard2To4, nil
	}
	return 0, flserr.Errorf("invalid high card. %s", highCard)
}
