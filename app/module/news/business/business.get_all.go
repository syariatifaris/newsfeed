package business

import (
	"context"

	"database/sql"

	"github.com/syariatifaris/arkeus/core/errors"
	"github.com/syariatifaris/arkeus/core/framework/business"
	"github.com/syariatifaris/arkeus/core/log/arklog"
	"github.com/syariatifaris/kumparan/app/module/news/model"
	newsrepo "github.com/syariatifaris/kumparan/app/module/news/repo"
)

const (
	statusAll = "all"
)

type getAllByStatusBusinessModelWithContext struct {
	newsRepo newsrepo.NewsRepository

	key string
	business.BusinessModel
}

func (bm *getAllByStatusBusinessModelWithContext) GetModelCtx(ctx context.Context) (interface{}, error) {
	if err := bm.Validate(nil); err != nil {
		arklog.ERROR.Println("GetModelCtx|", "validation error", err.Error())
		return nil, errors.New(errors.InvalidRequest)
	}

	var news []*model.News
	var err error

	if bm.key == statusAll {
		news, err = bm.newsRepo.GetAll(ctx)
	} else {
		news, err = bm.newsRepo.GetAllByStatus(ctx, bm.key)
	}

	if err != nil && err != sql.ErrNoRows {
		arklog.ERROR.Println("GetModelCtx|", "db error", err.Error())
		return nil, errors.New(errors.DatabaseExecutionFail)
	}

	return news, nil
}

func (bm *getAllByStatusBusinessModelWithContext) Validate(obj interface{}) error {
	bm.Validator().Required(bm.key).SetErrorMessage("news status is required")
	ok, errs := bm.Validator().Validate()

	if !ok && len(errs) > 0 {
		return errors.New(bm.ErrorsToString(errs))
	}

	return nil
}
