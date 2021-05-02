package connection

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/lib/pq"
)

// Store is a wrapper for GORM -- an ORM library.
type ConnectionPool struct {
	*gorm.DB
}

// Gorm is a connection pool that's safe for sharing between handlers/requests
// https://github.com/jinzhu/gorm/issues/246
var _connectionPool ConnectionPool

// connectionPool connects to the database using Gorm and sets logging mode.
func connectionPool(dialect string, dbString string, debugLogging bool) (db *gorm.DB) {
	db, err := gorm.Open(dialect, dbString)
	if err != nil {
		log.Panic("failed to connect database")
	}

	db.LogMode(debugLogging)
	return db
}

// Connect and setter for the data store. Parameters dialect and dbString should correspond with:
//  gorm.Open(dialect, dbString)
func Connect(dialect string, dbString string, debugLogging bool) ConnectionPool {
	_connectionPool = ConnectionPool{connectionPool(dialect, dbString, debugLogging)}
	return _connectionPool
}

// Get the active data store.
func Get() ConnectionPool {
	return _connectionPool
}
