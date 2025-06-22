package database

import (
	"database/sql"
	"errors"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type OrmDB struct {
	OrmInstance *gorm.DB
	Database    Database
}

func OpenOrmWithDatabase(database Database) (*OrmDB, error) {
	ormDB := OrmDB{}
	if database == nil {
		return nil, errors.New("database object is nil")
	}
	var err error
	d := database.Get().(*sql.DB)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: d,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	ormDB.OrmInstance = gormDB
	ormDB.Database = database
	return &ormDB, nil

}

func OpenOrm(host string, port int, user string, password string, dbname string) (*OrmDB, error) {

	postgresDb, err := OpenPostgresSqlDatabase(host, port, user, password, dbname)
	if err != nil {
		return nil, err
	}
	err = postgresDb.Ping()
	if err != nil {
		return nil, err
	}
	return OpenOrmWithDatabase(postgresDb)
}
