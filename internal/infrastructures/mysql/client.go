package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mariojuzar/fleet-inventory/config"
)

type DB struct {
	Conn *sqlx.DB
}

type Options struct {
	Cfg config.DBConfig
}

func New(opt *Options) (*DB, error) {
	db, err := sqlx.Open("mysql", opt.Cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(opt.Cfg.MaxIdleConn)
	db.SetMaxOpenConns(opt.Cfg.MaxConn)
	db.SetConnMaxLifetime(opt.Cfg.ConnMaxLifetime)

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}
