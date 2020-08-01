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
	position.ipost_flop_action_order
FROM
	player_position position
WHERE
	position_type = ?
`)

	e := playerposition.Entity{}
	err := c.db.QueryRowContext(
		ctx,
		text,
		position.String(),
	).Scan(
		e.ID,
		e.PositionType,
		e.PostFlopActionOrder,
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
	INNER JOIN cards AS middleCard
		ON middleCard.id = board.middle_card_id
	INNER JOIN card_numbers AS middleCardNum
		ON middleCardNum.id = middleCard.card_number_id
	INNER JOIN cards AS lowCard
		ON lowCard.id = board.low_card_id
	INNER JOIN card_numbers AS lowCardNum
		ON lowCardNum.id = lowCard.card_number_id
	INNER JOIN heads_up_situations AS headsup 
		ON headsup.id = situation.heads_up_situation_id
	INNER JOIN player_positions AS inPos
		ON inPos.id = headsup.in_position_id
	INNER JOIN player_positions AS outOfPos
		ON outOfPos.id = headsup.out_of_position_id
LIMIT
	%d
`, 10)

	q, err := c.db.QueryContext(ctx, text)
	if err != nil {
		return nil, err
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
