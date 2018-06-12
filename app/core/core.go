package core

import (
	"github.com/jmoiron/sqlx"
	fwhandler "github.com/syariatifaris/arkeus/core/framework/handler"
	"github.com/syariatifaris/arkeus/core/inject"
	"github.com/syariatifaris/arkeus/core/net"
	"github.com/syariatifaris/kumparan/app/core/config"
	"github.com/syariatifaris/kumparan/app/core/db"
	"github.com/syariatifaris/kumparan/app/handler"
	newsbusiness "github.com/syariatifaris/kumparan/app/module/news/business"
	newsrepo "github.com/syariatifaris/kumparan/app/module/news/repo"
	topicrepo "github.com/syariatifaris/kumparan/app/module/topic/repo"
)

//gets new dependencies
func NewDependencies() inject.Injection {
	var (
		cfg *config.ConfigurationData
		rdb *sqlx.DB

		router      net.Router
		newsFactory newsbusiness.Factory
		newsRepo    newsrepo.NewsRepository
		topicRepo   topicrepo.TopicRepository
		baseHandler *fwhandler.BaseHandler
		newsHandler *handler.NewsHandler
	)

	di := inject.NewDependencyInjection()
	di.AddDependency(&router, net.NewRouter)

	di.AddDependency(&cfg, config.NewConfiguration)
	di.Resolve(&cfg)

	di.AddDependency(&rdb, db.NewSqlxConnection, &cfg.Database)
	di.Resolve(&rdb)
	di.AddDependency(&newsRepo, newsrepo.NewNewsRepository, &rdb)
	di.AddDependency(&topicRepo, topicrepo.NewTopicRepository, &rdb)

	di.AddDependency(&baseHandler, fwhandler.NewSimpleBaseHandler)

	di.AddDependency(&newsFactory, newsbusiness.NewFactory, &newsRepo, &topicRepo)
	di.AddDependency(&newsHandler, handler.NewNewsHandler, &newsFactory)

	return di
}
