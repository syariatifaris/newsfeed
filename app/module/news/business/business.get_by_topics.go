package business

import (
	"context"

	"github.com/syariatifaris/arkeus/core/errors"
	"github.com/syariatifaris/arkeus/core/framework/business"
	"github.com/syariatifaris/arkeus/core/log/arklog"
	"github.com/syariatifaris/kumparan/app/module/news/model"
	"github.com/syariatifaris/kumparan/app/module/news/repo"
)

type getByTopicsBusinessModelWithContext struct {
	request model.GetByTopicsRequest
	repo    repo.NewsRepository
	business.BusinessModel
}

func (bm *getByTopicsBusinessModelWithContext) GetModelCtx(ctx context.Context) (interface{}, error) {
	if err := bm.Validate(nil); err != nil {
		arklog.ERROR.Println("GetModelCtx|", "validation error", err.Error())
		return nil, errors.New(errors.InvalidRequest)
	}

	return bm.repo.GetAllByTopics(ctx, bm.request.Topics)
}

func (bm *getByTopicsBusinessModelWithContext) Validate(obj interface{}) error {
	bm.Validator().MinSize(bm.request.Topics, 1).SetErrorMessage("invalid searched topic length")
	ok, errs := bm.Validator().Validate()

	if !ok && len(errs) > 0 {
		return errors.New(bm.ErrorsToString(errs))
	}

	return nil
}
