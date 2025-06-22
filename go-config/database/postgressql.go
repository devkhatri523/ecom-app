package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/devkhatri523/ecom-app/go-config/config"
)

type PostgresDatabase struct {
	db *sql.DB
}

func OpenPostgresSqlDatabase(host string, port int, userName string, password string, dbname string) (*PostgresDatabase, error) {

	postgresDb := &PostgresDatabase{
		nil,
	}
	connMaxLifetime := config.Default().GetInt("db.postgressql.connMaxLifetime")
	maxOpenConn := config.Default().GetInt("db.postgressql.maxOpenConn")
	maxIdleConn := config.Default().GetInt("db.postgressql.maxIdleConn")
	param := config.Default().GetString("db.postgressql.param")
	postgresDb.Open(Options{
		Host:            host,
		Port:            port,
		UserName:        userName,
		Password:        password,
		DatabaseName:    dbname,
		ConnMaxLifetime: time.Duration(connMaxLifetime),
		MaxOpenConns:    maxOpenConn,
		MaxIdleConns:    maxIdleConn,
		PARAM:           param,
	})

	return postgresDb, nil
}

func (m *PostgresDatabase) Open(options Options) {
	dbs, err := BuildDns(options)
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("Opening database  connection on host: %s port:%d database: %s username:%s", options.Host,
		options.Port, options.DatabaseName, options.UserName)
	db, err := sql.Open("postgres", dbs)
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("Configuring database  connection on host: %s port:%d database: %s username:%s", options.Host,
		options.Port, options.DatabaseName, options.UserName)
	db.SetConnMaxIdleTime(options.ConnMaxLifetime)
	db.SetMaxIdleConns(options.MaxIdleConns)
	db.SetMaxOpenConns(options.MaxOpenConns)
	m.db = db
}

func (m *PostgresDatabase) Close() {
	if m.db != nil {
		err := m.db.Close()
		if err != nil {
			fmt.Sprintf("Error while closing database %s", err)
			return
		}
	}
	fmt.Sprintf("Database connection closed")
}

func (m *PostgresDatabase) Get() interface{} {
	if m.db == nil {
		panic("Database connection is not initiated .Please call Open()")
	}
	return m.db
}
func (m *PostgresDatabase) Ping() error {
	if m.db == nil {
		panic("Database connection is not initiated .Please call Open()")
	}
	err := m.db.Ping()
	if err != nil {
		return err
	} else {
		fmt.Sprintf("Postgressql database connection succeeded")
	}
	return nil
}
