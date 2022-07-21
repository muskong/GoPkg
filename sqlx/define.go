package sqlx

import (
	"github.com/jmoiron/sqlx"
	"github.com/muskong/GoPkg/zaplog"
)

// BindNamed binds a query using the DB driver's bindvar type.
func BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return xdb.BindNamed(query, arg)
}

// NamedQuery using this DB.
// Any named placeholder parameters are replaced with fields from arg.
func NamedQuery(query string, arg interface{}) (*sqlx.Rows, error) {
	return xdb.NamedQuery(query, arg)
}

// NamedExec using this DB.
// Any named placeholder parameters are replaced with fields from arg.
func NamedInsert(query string, arg interface{}) int64 {
	result, err := xdb.NamedExec(query, arg)
	if err != nil {
		zaplog.Sugar.Error(err)
		return 0
	}

	insertId, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	return insertId
}
func NamedUpdate(query string, arg interface{}) int64 {
	result, err := xdb.NamedExec(query, arg)
	if err != nil {
		zaplog.Sugar.Error(err)
		return 0
	}

	rows, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	return rows
}

func Select(dest interface{}, query string, args ...interface{}) error {
	return xdb.Select(dest, query, args...)
}
func Get(dest interface{}, query string, args ...interface{}) error {
	return xdb.Get(dest, query, args...)
}

// MustBegin starts a transaction, and panics on error.  Returns an *sqlx.Tx instead
// of an *sql.Tx.
func MustBegin() *sqlx.Tx {
	tx, err := xdb.Beginx()
	if err != nil {
		zaplog.Sugar.Error(err)
	}
	return tx
}

// Queryx queries the database and returns an *sqlx.Rows.
// Any placeholder parameters are replaced with supplied args.
func Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return xdb.Queryx(query, args...)
}

// QueryRowx queries the database and returns an *sqlx.Row.
// Any placeholder parameters are replaced with supplied args.
func QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return xdb.QueryRowx(query, args...)
}

// MustExec (panic) runs MustExec using this database.
// Any placeholder parameters are replaced with supplied args.
func MustExec(query string, args ...interface{}) int64 {
	result := xdb.MustExec(query, args...)

	rows, err := result.RowsAffected() // 操作影响的行数
	if err != nil {
		zaplog.Sugar.Error(err)
	}

	return rows
}

// Preparex returns an sqlx.Stmt instead of a sql.Stmt
func Preparex(query string) (*sqlx.Stmt, error) {
	return xdb.Preparex(query)
}

// PrepareNamed returns an sqlx.NamedStmt
func PrepareNamed(query string) (*sqlx.NamedStmt, error) {
	return xdb.PrepareNamed(query)
}
