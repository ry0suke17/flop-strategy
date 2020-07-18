package flopsituationapp

import (
	"context"

	"github.com/yneee/flop-strategy/domain/flopsituationlist"
)

// ListFlopSituations はフロップシチュエーションのリストを取得する。
func (a *App) ListFlopSituations(
	ctx context.Context,
) ([]*flopsituationlist.Entity, error) {
	return a.db.ListFlopSituations(ctx)
}
