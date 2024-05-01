package postgresql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type PgDatabase struct {
	DBConfig Config
	Database *sql.DB
}

func (m *PgDatabase) GetDB() *sql.DB {
	return m.Database
}

func (md *PgDatabase) Open() error {
	return md.open()
}

func NewPgDatabase(config Config) *PgDatabase {
	md := PgDatabase{}
	md.DBConfig = config
	return &md
}

func (md *PgDatabase) open() error {
	fmt.Println("here")
	db, err := sql.Open("postgres", md.DBConfig.DSN())
	fmt.Printf("dsn : %v \n", md.DBConfig.DSN())
	if err != nil {
		panic(err)
	}
	fmt.Println("here2")
	//err = db.Ping()
	//if err != nil {
	//	panic(err)
	//}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	md.Database = db
	fmt.Println("here3")
	return nil
}
