package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/arkeus/core/errors"
	"github.com/syariatifaris/arkeus/core/log/arklog"
	"github.com/syariatifaris/kumparan/app/core/db"
	"github.com/syariatifaris/kumparan/app/module/topic/model"
)

func NewTopicRepository(db *sqlx.DB) TopicRepository {
	return &topicRepository{
		db: db,
	}
}

type TxFunc func(ctx context.Context, tx *sqlx.Tx) error

type TopicRepository interface {
	GetIDsByNames(ctx context.Context, names []string) ([]int64, error)
	AddNewsRelation(context.Context, *sqlx.Tx, *model.TopicNews) error
	ExecuteInTx(ctx context.Context, f TxFunc) error
}

type topicRepository struct {
	db *sqlx.DB
}

func (t *topicRepository) GetIDsByNames(ctx context.Context, names []string) ([]int64, error) {
	var ids []int64
	query, args, err := db.RebindQuery(getIDsByNames, names)
	if err != nil {
		arklog.ERROR.Println("TopicRepo|GetIDsByNames", "db error:", err.Error())
		return nil, errors.New(errors.DatabaseExecutionFail)
	}

	query = t.db.Rebind(query)
	err = t.db.SelectContext(ctx, &ids, query, args...)
	if err != nil {
		arklog.ERROR.Println("TopicRepo|GetIDsByNames", "db error:", err.Error())
		return nil, errors.New(errors.DatabaseExecutionFail)
	}

	return ids, nil
}

func (t *topicRepository) AddNewsRelation(ctx context.Context, tx *sqlx.Tx, topicNews *model.TopicNews) error {
	query, args, err := db.RebindQuery(insertTopicNews, topicNews.TopicID, topicNews.NewsID)
	if err != nil {
		return err
	}

	_, err = tx.QueryContext(ctx, t.db.Rebind(query), args...)
	if err != nil {
		return err
	}

	return nil
}

func (t *topicRepository) ExecuteInTx(ctx context.Context, txf TxFunc) error {
	if ctx == nil {
		ctx = context.Background()
	}
	tx, err := t.db.Beginx()
	if err != nil {
		return err
	}

	err = db.ExecuteInTx(ctx, tx, func() error {
		err := txf(ctx, tx)
		return err
	})

	return err
}
