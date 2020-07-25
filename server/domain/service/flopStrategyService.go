package service

import (
	"context"

	"github.com/yneee/flop-strategy/domain/model/flopsituationlist"
)

// DBRepository は DB のリポジトリを表す。
type DBRepository interface {
	// ListFlopSituations はフロップシチュエーションのリストを取得する。
	ListFlopSituations(
		ctx context.Context,
	) ([]*flopsituationlist.Entity, error)
}

// FlopStrtategyService はフロップシチュエーションを扱うサービスを表す。
type FlopStrtategyService struct {
	db DBRepository
}

// NewFlopStrtategyService は新しいサービスを生成する。
func NewFlopStrtategyService(
	db DBRepository,
) *FlopStrtategyService {
	return &FlopStrtategyService{
		db: db,
	}
}
