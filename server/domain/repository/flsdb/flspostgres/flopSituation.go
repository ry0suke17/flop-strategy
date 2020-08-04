package flspostgres

import (
	"context"
	"fmt"

	"github.com/yneee/flop-strategy/domain/model/board"
	"github.com/yneee/flop-strategy/domain/model/flopsituationlist"
	"github.com/yneee/flop-strategy/domain/model/playerposition"
	"github.com/yneee/flop-strategy/infra/flserr"
)

// GetPlayerPostion はプレイヤーのポジションを取得する。
func (c *Client) GetPlayerPostion(
	ctx context.Context,
	position playerposition.Position,
) (*playerposition.Entity, error) {
	text := fmt.Sprint(`
SELECT
	position.id,
	position.position_type,
	position.post_flop_action_order
FROM
	player_positions position
WHERE
	position_type = $1
`)

	e := playerposition.Entity{}
	err := c.db.QueryRowContext(
		ctx,
		text,
		position.String(),
	).Scan(
		&e.ID,
		&e.PositionType,
		&e.PostFlopActionOrder,
	)
	if err != nil {
		return nil, err
	}
	return &e, nil
}

// ListFlopSituations はフロップシチュエーションのリストを取得する。
func (c *Client) ListFlopSituations(
	ctx context.Context,
	inPosition playerposition.Position,
	outOfPosition playerposition.Position,
	potType board.PotType,
	highCard board.HighCard,
	boardPairedType board.PairedType,
	boardSuitsType board.SuitsType,
) ([]*flopsituationlist.Entity, error) {
	highCardFilterText, err := highCardFilterText(highCard)
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	boardPairedNumber, err := boardPairedNumber(boardPairedType)
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	boardSuitsNumber, err := boardSuitsNumber(boardSuitsType)
	if err != nil {
		return nil, flserr.Wrap(err)
	}

	// TODO: connected とそうでないのとで分ける必要がある
	text := fmt.Sprintf(`
SELECT
	param.in_position_bet_frequency,
	param.out_of_position_bet_frequency,
	param.in_position_check_frequency,
	param.out_of_position_check_frequency,
	param.in_position_33_bet_frequency,
	param.out_of_position_33_bet_frequency,
	param.in_position_67_bet_frequency,
	param.out_of_position_67_bet_frequency,
	param.in_position_equity,
	param.out_of_position_equity
FROM
	flop_situations situation
	INNER JOIN flop_situation_parameters AS param
		ON param.flop_situation_id = situation.id
	LEFT JOIN flop_situation_images AS image
		ON image.flop_situation_id = situation.id
	INNER JOIN boards AS board
		ON board.id = situation.board_id
	INNER JOIN cards AS highCard
		ON highCard.id = board.high_card_id
	INNER JOIN card_numbers AS highCardNum
		ON highCardNum.id = highCard.card_number_id
	INNER JOIN card_suits AS highCardSuit
		ON highCardSuit.id = highCard.card_suit_id
	INNER JOIN cards AS middleCard
		ON middleCard.id = board.middle_card_id
	INNER JOIN card_numbers AS middleCardNum
		ON middleCardNum.id = middleCard.card_number_id
	INNER JOIN card_suits AS middleCardSuit
		ON middleCardSuit.id = middleCard.card_suit_id
	INNER JOIN cards AS lowCard
		ON lowCard.id = board.low_card_id
	INNER JOIN card_numbers AS lowCardNum
		ON lowCardNum.id = lowCard.card_number_id
	INNER JOIN card_suits AS lowCardSuit
		ON lowCardSuit.id = lowCard.card_suit_id
	INNER JOIN heads_up_situations AS headsup 
		ON headsup.id = situation.heads_up_situation_id
	INNER JOIN player_positions AS inPos
		ON inPos.id = headsup.in_position_id
	INNER JOIN player_positions AS outOfPos
		ON outOfPos.id = headsup.out_of_position_id
WHERE
		inPos.position_type = $1
	AND	outOfPos.position_type = $2
	AND	headsup.pot_type = $3
	AND (
		SELECT
			array_length(array_agg(DISTINCT value), 1)
		FROM
			unnest(ARRAY[highCardNum.value, middleCardNum.value, lowCardNum.value]) AS value
	) = $4
	AND (
		SELECT
			array_length(array_agg(DISTINCT value), 1)
		FROM
			unnest(ARRAY[highCardSuit.value, middleCardSuit.value, lowCardSuit.value]) AS value
	) = $5
	AND %s
LIMIT
	%d
`, highCardFilterText, 100)

	q, err := c.db.QueryContext(
		ctx,
		text,
		inPosition.String(),
		outOfPosition.String(),
		potType.String(),
		boardPairedNumber,
		boardSuitsNumber,
	)
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	defer func() {
		_ = q.Close()
	}()

	var list []*flopsituationlist.Entity
	for q.Next() {
		a := flopsituationlist.Entity{}
		err = q.Scan(
			&a.InPositionBetFrequency,
			&a.OutOfPositionBetFrequency,
			&a.InPositionCheckFrequency,
			&a.OutOfPositionCheckFrequency,
			&a.InPosition33BetFrequency,
			&a.OutOfPosition33BetFrequency,
			&a.InPosition67BetFrequency,
			&a.OutOfPosition67BetFrequency,
			&a.InPositionEquity,
			&a.OutOfPositionEquity,
		)
		if err != nil {
			return nil, flserr.Wrap(err)
		}
		list = append(list, &a)
	}

	return list, nil
}

func highCardFilterText(highCard board.HighCard) (string, error) {
	switch highCard {
	case board.HighCardA:
		return "highCardNum.value = 14", nil
	case board.HighCardK:
		return "highCardNum.value = 13", nil
	case board.HighCardQ:
		return "highCardNum.value = 12", nil
	case board.HighCardJ:
		return "highCardNum.value = 11", nil
	case board.HighCardT:
		return "highCardNum.value = 10", nil
	case board.HighCard8To9:
		return "highCardNum.value BETWEEN 8 AND 9", nil
	case board.HighCard5To7:
		return "highCardNum.value BETWEEN 5 AND 7", nil
	case board.HighCard2To4:
		return "highCardNum.value BETWEEN 2 AND 4", nil
	}
	return "", flserr.Errorf("invalid highCard. highCard=%d", highCard)
}

func boardPairedNumber(boardPairedType board.PairedType) (int, error) {
	switch boardPairedType {
	case board.PairedTypeUnpaired:
		return 3, nil
	case board.PairedTypePaired:
		return 2, nil
	case board.PairedTypeTrips:
		return 1, nil
	}
	return 0, flserr.Errorf("invalid boardPairedType. boardPairedType=%d", boardPairedType)
}

func boardSuitsNumber(boardSuitsType board.SuitsType) (int, error) {
	switch boardSuitsType {
	case board.SuitsTypeMonoTone:
		return 1, nil
	case board.SuitsTypeTwoTone:
		return 2, nil
	case board.SuitsTypeRainbow:
		return 3, nil
	}
	return 0, flserr.Errorf("invalid boardSuitsType. boardSuitsType=%d", boardSuitsType)
}
