package server

import (
	"log"
	"net/http"
	"strings"

	"github.com/yneee/flop-strategy/infra/flserr"

	"github.com/yneee/flop-strategy/transport/openapi"

	"github.com/yneee/flop-strategy/domain/service"
	"github.com/yneee/flop-strategy/transport/openapi/api"
)

// FlopStrategyServer はフロップ戦略のサーバーを表す。
type FlopStrategyServer struct {
	service *service.FlopStrtategyService
}

// NewFlopStrategyServer は新しいサーバーを返す。
func NewFlopStrategyServer(s *service.FlopStrtategyService) api.Router {
	return &FlopStrategyServer{
		service: s,
	}
}

// Routes は OpenAPI の Routes の実装を表す。
func (s *FlopStrategyServer) Routes() api.Routes {
	return api.Routes{
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
	resp, err := s.getFlopSituationsParameter(r)
	if err != nil {
		errorHandling(w, err)
		return
	}

	err = api.EncodeJSONResponse(resp, nil, w)
	if err != nil {
		errorHandling(w, err)
		return
	}
}

func errorHandling(w http.ResponseWriter, err error) {
	status := http.StatusBadRequest
	respErr := &api.Error{
		// TODO: とりあえず空文字にしている。
		// クライアント側で厳密にメッセージを分ける場合はここで Code を指定すると良さそう。
		Code:             "",
		Message:          err.Error(),
		LocalizedMessage: "",
	}
	encodeErr := api.EncodeJSONResponse(respErr, &status, w)
	if encodeErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed response encode. err=%v\n", encodeErr)
	}
}

func (s *FlopStrategyServer) getFlopSituationsParameter(r *http.Request) (*api.GetFlopSituationsParameterResponse, error) {
	query := r.URL.Query()
	heroPosition, err := openapi.PlayerPostion(api.PlayerPosition(query.Get("hero-position")))
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	villainPosition, err := openapi.PlayerPostion(api.PlayerPosition(query.Get("villain-position")))
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	potType, err := openapi.BoardPotType(api.PotType(query.Get("pot-type")))
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	highCard, err := openapi.BoardHighCard(api.HighCard(query.Get("high-card")))
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	boardPairType, err := openapi.BoardPairType(api.BoardPairType(query.Get("board-pair-type")))
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	boardSuitsType, err := openapi.BoardSuitsType(api.BoardSuitsType(query.Get("board-suits-type")))
	if err != nil {
		return nil, flserr.Wrap(err)
	}
	boardConnectType, err := openapi.BoardConnectType(api.BoardConnectType(query.Get("board-connect-type")))
	if err != nil {
		return nil, flserr.Wrap(err)
	}

	list, heroPositionType, err := s.service.GetFlopSituationsParameter(
		r.Context(),
		heroPosition,
		villainPosition,
		potType,
		highCard,
		boardPairType,
		boardSuitsType,
		boardConnectType,
	)
	if err != nil {
		return nil, flserr.Wrap(err)
	}

	heroPosType, err := openapi.ToAPIPlayerPosition(heroPositionType)
	if err != nil {
		return nil, flserr.Wrap(err)
	}

	return &api.GetFlopSituationsParameterResponse{
		IpBetFreq:        list.AvgInPositionBetFrequency(),
		OopBetFreq:       list.AvgOutOfPositionBetFrequency(),
		IpCheckFreq:      list.AvgInPositionCheckFrequency(),
		OopCheckFreq:     list.AvgOutOfPositionCheckFrequency(),
		Ip33BetFreq:      list.AvgInPosition33BetFrequency(),
		Oop33BetFreq:     list.AvgOutOfPosition33BetFrequency(),
		Ip67BetFreq:      list.AvgInPosition67BetFrequency(),
		Oop67BetFreq:     list.AvgOutOfPosition67BetFrequency(),
		IpEquity:         list.AvgInPositionEquity(),
		OopEquity:        list.AvgOutOfPositionEquity(),
		HeroPositionType: heroPosType,
	}, nil
}
