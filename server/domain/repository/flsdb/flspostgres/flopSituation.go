package flspostgres

import (
	"context"
	"fmt"

	"github.com/yneee/flop-strategy/domain/model/flopsituationlist"
	"github.com/yneee/flop-strategy/infra/flserr"
)

// ListFlopSituations はフロップシチュエーションのリストを取得する。
func (c *Client) ListFlopSituations(
	ctx context.Context,
) ([]*flopsituationlist.Entity, error) {
	text := fmt.Sprintf(`
SELECT
	param.in_position_bet_frequency,
	param.out_of_position_bet_frequency
FROM
	flop_situations situation
	INNER JOIN flop_situation_parameters AS param
		ON param.flop_situation_id = situation.id
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
		err = q.Scan(&a.InPositionBetFrequency, &a.OutOfPositionBetFrequency)
		if err != nil {
			return nil, flserr.Wrap(err)
		}
		list = append(list, &a)
	}

	return list, nil
}
