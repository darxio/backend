package connection

import (
	"sync"

	"github.com/jackc/pgx"
)

var DB *pgx.ConnPool
var once sync.Once

const maxConn = 8

func Connect() (conn *pgx.ConnPool) {
	once.Do(func() {
		connConfig := pgx.ConnConfig{
			User:     "ksu",
			Host:     "localhost",
			Port:     5432,
			Database: "backend_db",
		}

		DB, _ = pgx.NewConnPool(pgx.ConnPoolConfig{
			ConnConfig:     connConfig,
			MaxConnections: maxConn,
		})
	})

	return DB
}
