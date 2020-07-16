package flopsituationapp

import (
	"context"

	"github.com/yneee/flop-strategy/domain/flopsituationlist"
)

// DBRepository は DB のリポジトリを表す。
type DBRepository interface {
	// ListFlopSituations はフロップシチュエーションのリストを取得する。
	ListFlopSituations(
		ctx context.Context,
	) ([]*flopsituationlist.Entity, error)
}

// App はフロップシチュエーションを管理するアプリケーションを表す。
type App struct {
	db DBRepository
}

// NewApp は新しいアプリケーションを生成する。
func NewApp(
	db DBRepository,
) *App {
	a := &App{
		db: db,
	}
	return a
}
