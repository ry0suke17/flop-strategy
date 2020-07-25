package server

import (
	"net/http"
	"strings"

	"github.com/yneee/flop-strategy/domain/service"

	"github.com/yneee/flop-strategy/infra/flsopenapi"
)

// FlopStrategyServer はフロップ戦略のサーバーを表す。
type FlopStrategyServer struct {
	service *service.FlopStrtategyService
}

// NewFlopStrategyServer は新しいサーバーを返す。
func NewFlopStrategyServer(s *service.FlopStrtategyService) flsopenapi.Router {
	return &FlopStrategyServer{
		service: s,
	}
}

// Routes は OpenAPI の Routes の実装を表す。
func (s *FlopStrategyServer) Routes() flsopenapi.Routes {
	return flsopenapi.Routes{
		{
			Name:        "GetFlopSituationsParameter",
			Method:      strings.ToUpper("Get"),
			Pattern:     "/flop/situations/parameter",
			HandlerFunc: s.GetFlopSituationsParameter,
		},
	}
}

// GetFlopSituationsParameter はフロップシチュエーションのパラメータを取得する
func (s *FlopStrategyServer) GetFlopSituationsParameter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	heroPosition := query.Get("heroPosition")
	villainPosition := query.Get("villainPosition")
	potType := query.Get("potType")
	highCard := query.Get("highCard")
	boardPairedType := query.Get("boardPairedType")
	boardSuitsType := query.Get("boardSuitsType")

	result, err := s.service.GetFlopSituationsParameter(
		r.Context(),
		heroPosition,
		villainPosition,
		potType,
		highCard,
		boardPairedType,
		boardSuitsType,
	)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	err = flsopenapi.EncodeJSONResponse(result, nil, w)
	if err != nil {
		w.WriteHeader(500)
		return
	}
}
