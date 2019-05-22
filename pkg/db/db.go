package db

import (
	"fmt"

	"github.com/dgraph-io/badger"
	"github.com/pkg/errors"
)

type DB struct {
	db *badger.DB
}

func NewDB(dbPath string) (*DB, error) {
	var dbOptions = badger.DefaultOptions
	dbOptions.Dir = dbPath
	dbOptions.ValueDir = dbPath
	var badgerDB, errOpenDB = badger.Open(dbOptions)
	if errOpenDB != nil {
		var msg = fmt.Sprintf("unable to open db in %q", dbPath)
		return nil, errors.Wrap(errOpenDB, msg)
	}
	return &DB{
		db: badgerDB,
	}, nil
}
