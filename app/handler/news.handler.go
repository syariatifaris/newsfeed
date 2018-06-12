package handler

import (
	"net/http"

	"github.com/syariatifaris/arkeus/core/errors"
	"github.com/syariatifaris/arkeus/core/framework/handler"
	"github.com/syariatifaris/arkeus/core/log/arklog"
	"github.com/syariatifaris/arkeus/core/net"
	"github.com/syariatifaris/kumparan/app/module/news/business"
	"github.com/syariatifaris/kumparan/app/module/news/model"
)

func NewNewsHandler(factory business.Factory) *NewsHandler {
	return &NewsHandler{
		factory: factory,
	}
}

type NewsHandler struct {
	factory business.Factory
	handler.BaseHandler
}

func (*NewsHandler) Name() string {
	return "NewsHandler"
}

func (s *NewsHandler) RegisterHandlers(router net.Router) {
	newsRouter := router.PathPrefix("/news").Subrouter()
	newsRouter.HandleFunc("/gets/{status}", s.NoAuthenticate(s.Index)).Methods(http.MethodGet)
	newsRouter.HandleFunc("/gets/topics", s.NoAuthenticate(s.GetByTopics)).Methods(http.MethodPost)
	newsRouter.HandleFunc("/add", s.NoAuthenticate(s.Add)).Methods(http.MethodPost)
}

func (s *NewsHandler) Index(r *http.Request) (interface{}, error) {
	status := s.GetQueryData(r, "status")
	bm := s.factory.GetAllByStatusBusinessModelCtx(status)
	return bm.GetModelCtx(r.Context())
}

func (s *NewsHandler) Add(r *http.Request) (interface{}, error) {
	var request model.AddNewsRequest
	err := s.GetPostData(r, &request)
	if err != nil {
		arklog.ERROR.Println("NewsHandler|Add|", "unable to obtain request", err.Error())
		return nil, errors.New(errors.InvalidRequest)
	}

	bm := s.factory.AddBusinessModelCtx(request)
	return bm.GetModelCtx(r.Context())
}

func (s *NewsHandler) GetByTopics(r *http.Request) (interface{}, error) {
	var request model.GetByTopicsRequest
	err := s.GetPostData(r, &request)
	if err != nil {
		arklog.ERROR.Println("NewsHandler|GetByTopics|", "unable to obtain request", err.Error())
		return nil, errors.New(errors.InvalidRequest)
	}

	bm := s.factory.GetAllByTopicsBusinessModelCtx(request)
	return bm.GetModelCtx(r.Context())
}
