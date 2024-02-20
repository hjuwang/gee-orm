package geeorm

import (
	"database/sql"
	"geeorm/log"
	"geeorm/session"
)

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {

	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	//ping 测试连接数据库是否正常
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (engine *Engine) Close() {
	err := engine.db.Close()
	if err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (engine Engine) NewSession() *session.Session {

	return session.New(engine.db)
}
