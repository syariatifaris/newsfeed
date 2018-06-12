package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/arkeus/core/errors"
	"github.com/syariatifaris/arkeus/core/log/arklog"
	"github.com/syariatifaris/kumparan/app/core/db"
	"github.com/syariatifaris/kumparan/app/module/news/model"
)

func NewNewsRepository(db *sqlx.DB) NewsRepository {
	return &newsRepository{
		db: db,
	}
}

type TxFunc func(ctx context.Context, tx *sqlx.Tx) error

type NewsRepository interface {
	GetAll(ctx context.Context) ([]*model.News, error)
	GetAllByStatus(ctx context.Context, key string) ([]*model.News, error)
	GetAllByTopics(ctx context.Context, topics []string) ([]*model.News, error)
	GetByID(ctx context.Context, id int64) *model.News
	Add(ctx context.Context, tx *sqlx.Tx, news *model.News) error
	ExecuteInTx(ctx context.Context, f TxFunc) error
}

type newsRepository struct {
	db *sqlx.DB
}

func (nr *newsRepository) ExecuteInTx(ctx context.Context, txf TxFunc) error {
	if ctx == nil {
		ctx = context.Background()
	}
	tx, err := nr.db.Beginx()
	if err != nil {
		return err
	}

	err = db.ExecuteInTx(ctx, tx, func() error {
		err := txf(ctx, tx)
		return err
	})

	return err
}

func (nr *newsRepository) GetAll(ctx context.Context) ([]*model.News, error) {
	var news []*model.News

	err := nr.db.SelectContext(ctx, &news, getAllNews)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (nr *newsRepository) GetAllByStatus(ctx context.Context, key string) ([]*model.News, error) {
	var news []*model.News

	err := nr.db.SelectContext(ctx, &news, getAllNewsByStatus, key)
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (nr *newsRepository) GetAllByTopics(ctx context.Context, topics []string) ([]*model.News, error) {
	var news []*model.News
	query, args, err := db.RebindQuery(getAllNewsByTopics, topics)
	if err != nil {
		arklog.ERROR.Println("NewsRepo|GetAllByTopics", "db error:", err.Error())
		return nil, errors.New(errors.DatabaseExecutionFail)
	}

	query = nr.db.Rebind(query)
	err = nr.db.SelectContext(ctx, &news, query, args...)
	if err != nil {
		arklog.ERROR.Println("NewsRepo|GetAllByTopics", "db error:", err.Error())
		return nil, errors.New(errors.DatabaseExecutionFail)
	}

	return news, nil
}

func (*newsRepository) GetByID(ctx context.Context, id int64) *model.News {
	return &model.News{}
}

func (nr *newsRepository) Add(ctx context.Context, tx *sqlx.Tx, news *model.News) error {
	query, args, err := db.RebindQuery(insertNews, news.Title, news.HtmlContent, news.Status)
	if err != nil {
		return err
	}

	var id int64
	err = tx.QueryRowContext(ctx, nr.db.Rebind(query), args...).Scan(&id)
	if err != nil {
		return err
	}

	news.NewsID = id
	return nil
}
