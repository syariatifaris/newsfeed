package business

import (
	"github.com/syariatifaris/arkeus/core/framework/business"
	"github.com/syariatifaris/kumparan/app/module/news/model"
	newsrepo "github.com/syariatifaris/kumparan/app/module/news/repo"
	topicrepo "github.com/syariatifaris/kumparan/app/module/topic/repo"
)

type Factory interface {
	GetAllByStatusBusinessModelCtx(key string) business.BaseBusinessModelWithContext
	GetAllByTopicsBusinessModelCtx(request model.GetByTopicsRequest) business.BaseBusinessModelWithContext
	GetByIDBusinessModelCtx() business.BaseBusinessModelWithContext
	AddBusinessModelCtx(request model.AddNewsRequest) business.BaseBusinessModelWithContext
}

func NewFactory(newsRepo newsrepo.NewsRepository, topicRepo topicrepo.TopicRepository) Factory {
	return &factoryImpl{
		newsRepo:  newsRepo,
		topicRepo: topicRepo,
	}
}

type factoryImpl struct {
	newsRepo  newsrepo.NewsRepository
	topicRepo topicrepo.TopicRepository
}

func (factory *factoryImpl) GetAllByStatusBusinessModelCtx(key string) business.BaseBusinessModelWithContext {
	return &getAllByStatusBusinessModelWithContext{
		newsRepo: factory.newsRepo,
		key:      key,
	}
}

func (factory *factoryImpl) GetAllByTopicsBusinessModelCtx(request model.GetByTopicsRequest) business.BaseBusinessModelWithContext {
	return &getByTopicsBusinessModelWithContext{
		request: request,
		repo:    factory.newsRepo,
	}
}

func (*factoryImpl) GetByIDBusinessModelCtx() business.BaseBusinessModelWithContext {
	panic("implement me")
}

func (factory *factoryImpl) AddBusinessModelCtx(request model.AddNewsRequest) business.BaseBusinessModelWithContext {
	return &addBusinessModelWithContext{
		request:   &request,
		newsRepo:  factory.newsRepo,
		topicRepo: factory.topicRepo,
	}
}
