package db

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/syariatifaris/arkeus/core/log/arklog"
	"github.com/syariatifaris/kumparan/app/core/config"
)

const dcsFormatPG = `user=%s password=%s host=%s dbname=%s sslmode=disable`

//NsqSqlxConnection createa new sqlx db connection
func NewSqlxConnection(cfg config.Database) *sqlx.DB {
	connectionString := fmt.Sprintf(dcsFormatPG,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name,
	)

	db, err := sqlx.Connect(cfg.Type, connectionString)
	if err != nil {
		arklog.ERROR.Panic("cannot create db connection", err.Error())
	}

	arklog.DEBUG.Println("connected to db", cfg.Name, "server", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	return db
}

//RebindQuery query IN translate
func RebindQuery(queryS string, param ...interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.In(queryS, param...)
	if err != nil {
		return "", nil, err
	}

	return query, args, nil
}

//ExecuteInTx executes in transaction
func ExecuteInTx(ctx context.Context, tx *sqlx.Tx, fn func() error) (err error) {
	defer func() {
		if err == nil {
			_ = tx.Commit()
		} else {
			_ = tx.Rollback()
		}
	}()

	err = fn()
	return err
}
