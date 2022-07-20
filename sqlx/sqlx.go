package sqlx

import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/muskong/GoService/pkg/zaplog"
)

var (
	dbOnce sync.Once
	xdb    *sqlx.DB
)

type SqlxConfig struct {
	Driver  string
	Dsn     string
	MaxOpen int
	MaxIdle int
}

func Init(xcfg *SqlxConfig) {
	dbOnce.Do(xcfg.initDbConnection)
}

func DbClose() {
	if err := xdb.Close(); err != nil {
		zaplog.Sugar.Fatal(err)
	}
}

func (x *SqlxConfig) initDbConnection() {

	var err error
	// Connect to MongoDB
	xdb = sqlx.MustConnect(x.Driver, x.Dsn)
	xdb.SetMaxOpenConns(x.MaxOpen)
	xdb.SetMaxIdleConns(x.MaxIdle)

	// Check the connection
	err = xdb.Ping()

	if err != nil {
		zaplog.Sugar.Panic(err)
	}
	return
}
