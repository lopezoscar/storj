package satellitedb

import "storj.io/pkg/satellite"

import "errors"

type DB struct {
	db *dbx.DB
}

func New(engine, paramstr string) (satellite.DB, error) {
	db, err := dbx.Open(engine, paramstr)
	return DB{db}, err
}

func (db *DB) Migrate() error {
	return errors.New("unimplemented")
}

func (db *DB) Users() satellite.Users {
	return Users{db}
}
