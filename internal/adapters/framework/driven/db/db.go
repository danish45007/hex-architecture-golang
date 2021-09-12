package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	//connect
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("Failed to connect to database %v", err)
		return nil, err
	}
	//test db
	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to ping db %v", err)
		return nil, err
	}
	return &Adapter{
		db: db,
	}, nil
}

//close database connection
func (da Adapter) CloseDbConn() {
	err := da.db.Close()
	if err != nil {
		log.Fatalf("Failed to close database connection %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {

	queryString, args, err := sq.
		Insert("arith_history").Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation).
		ToSql()
	if err != nil {
		log.Fatalf("Failed to form SQL query %v", err)
		return err
	}

	_, err = da.db.Exec(queryString, args...)
	if err != nil {
		log.Fatalf("Failed to execute SQL query %v", err)
		return err
	}

	return nil
}
