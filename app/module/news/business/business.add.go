package business

import (
	"context"

	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/arkeus/core/errors"
	"github.com/syariatifaris/arkeus/core/framework/business"
	"github.com/syariatifaris/arkeus/core/log/arklog"
	"github.com/syariatifaris/kumparan/app/module/news/model"
	newsrepo "github.com/syariatifaris/kumparan/app/module/news/repo"
	topicmodel "github.com/syariatifaris/kumparan/app/module/topic/model"
	topicrepo "github.com/syariatifaris/kumparan/app/module/topic/repo"
)

type addBusinessModelWithContext struct {
	newsRepo  newsrepo.NewsRepository
	topicRepo topicrepo.TopicRepository

	request *model.AddNewsRequest
	business.BusinessModel
}

func (bm *addBusinessModelWithContext) GetModelCtx(ctx context.Context) (interface{}, error) {
	if err := bm.Validate(nil); err != nil {
		arklog.ERROR.Println("GetModelCtx|", "validation error:", err.Error())
		return nil, errors.New(errors.InvalidRequest)
	}

	topicIDs, err := bm.topicRepo.GetIDsByNames(ctx, bm.request.Topics)
	if err != nil {
		if err == sql.ErrNoRows {
			arklog.ERROR.Print("GetModelCtx|", "db error:", err.Error())
			return nil, errors.New(errors.InvalidRequest)
		}

		return nil, errors.New(errors.DatabaseExecutionFail)
	}

	if len(topicIDs) == 0 {
		arklog.ERROR.Print("GetModelCtx|", "topic not found:", err.Error())
		return nil, errors.New(errors.InvalidRequest)
	}

	arklog.DEBUG.Println("GetModelCtx|", topicIDs)

	err = bm.newsRepo.ExecuteInTx(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
		return bm.newsRepo.Add(ctx, tx, bm.request.News)
	})

	if err != nil {
		arklog.ERROR.Print("GetModelCtx|", "db error:", err.Error())
		return nil, errors.New(errors.DatabaseExecutionFail)
	}

	for _, id := range topicIDs {
		err := bm.topicRepo.ExecuteInTx(ctx, func(ctx context.Context, tx *sqlx.Tx) error {
			err = bm.topicRepo.AddNewsRelation(ctx, tx, &topicmodel.TopicNews{
				NewsID:  bm.request.News.NewsID,
				TopicID: id,
			})

			if err != nil {
				return err
			}

			return nil
		})
		if err != nil {
			arklog.ERROR.Print("GetModelCtx|", "db error:", err.Error())
			return nil, errors.New(errors.DatabaseExecutionFail)
		}
	}

	return bm.request, nil
}

func (bm *addBusinessModelWithContext) Validate(obj interface{}) error {
	bm.Validator().Required(bm.request.News.Title).SetErrorMessage("title is required")
	bm.Validator().MinSize(bm.request.News.Title, 10).SetErrorMessage("title should have minimum 10 character")
	bm.Validator().Required(bm.request.News.HtmlContent).SetErrorMessage("content is required")
	bm.Validator().MinSize(bm.request.Topics, 1).SetErrorMessage("news should have at least 1 topic")

	ok, errs := bm.Validator().Validate()

	if !ok && len(errs) > 0 {
		return errors.New(bm.ErrorsToString(errs))
	}

	return nil
}
