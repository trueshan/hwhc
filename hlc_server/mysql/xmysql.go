package mysql

import (
	"database/sql"

	"github.com/hwhc/hlc_server/log"
)

type XMySQL struct {
	db          *sql.DB
	Transaction bool
	tx          *sql.Tx
	Finished    bool
}

func (mysql MySQL) NewXMySQL(tx bool) *XMySQL {
	if tx {
		transaction, err := mysql.database.Begin()
		if err != nil {
			log.Error("Begin transaction failed, return false, %v", err)
			return nil
		}
		return &XMySQL{
			db:          mysql.database,
			Transaction: true,
			Finished:    false,
			tx:          transaction,
		}
	}
	return &XMySQL{
		db:          mysql.database,
		Transaction: false,
		tx:          nil,
	}
}

func (xmysql *XMySQL) Exec(query string, args ...interface{}) (sql.Result, error) {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.Exec(query, args...)
	} else {
		return xmysql.db.Exec(query, args...)
	}
}

func (xmysql *XMySQL) QueryRow(query string, args ...interface{}) *sql.Row {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.QueryRow(query, args...)
	} else {
		return xmysql.db.QueryRow(query, args...)
	}
}

func (xmysql *XMySQL) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if xmysql.Transaction && !xmysql.Finished {
		return xmysql.tx.Query(query, args...)
	} else {
		return xmysql.db.Query(query, args...)
	}
}

func (xmysql *XMySQL) Commit() {
	if xmysql.Transaction && !xmysql.Finished {
		err := xmysql.tx.Commit()
		if err != nil {
			log.Error("There is an error when commit transaction")
		}
		xmysql.Finished = true
	}
}

func (xmysql *XMySQL) Rollback() {
	if xmysql.Transaction && !xmysql.Finished {
		err := xmysql.tx.Rollback()
		if err != nil {
			log.Error("There is an error when rooback transaction")
		}
		xmysql.Finished = true
	}
}
