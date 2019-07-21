// package gdbc is created to represents low level database interfaces in order to have an unified way to
// access database handler.
// It is created to make it easier to handle certain database operations like transactions, database factory.
// It is ony a POC, not a mature solution
package gdbc

import (
	"context"
	"database/sql"
	"github.com/go-kivik/kivik"
)

// SqlGdbc (SQL Go database connection) is a wrapper for SQL database handler ( can be *sql.DB or *sql.Tx)
// It should be able to work with all SQL data that follows SQL standard.
type SqlGdbc interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	// If need transaction support, add this interface
	Transactioner
}

// NoSqlGdbc (NoSQL Go database connection) is a wrapper for NoSql database handler.
// Currently, it is just a POC instead of a mature implementation.
// To make it real, you need to remove kivik related types and need to do conversion between generic types and
// individual NoSql types
// It also doesn't includes all methods, just the one I need. You can add more when there is a need.
type NoSqlGdbc interface {
	// The method name of underline database was Query(), but since it conflicts with the name with Query() in SqlGdbc,
	// so have to change to a different name
	QueryNoSql(ctx context.Context, ddoc string, view string) (*kivik.Rows, error)
	Put(ctx context.Context, docID string, doc interface{}, options ...kivik.Options) (rev string, err error)
	Get(ctx context.Context, docID string, options ...kivik.Options) (*kivik.Row, error)
	Find(ctx context.Context, query interface{}) (*kivik.Rows, error)
	AllDocs(ctx context.Context, options ...kivik.Options) (*kivik.Rows, error)
}

// gdbc is an unified way to handle database connections. As long as a data service type implements this interface,
// it can work with both SQL and NOSQL underline database.
// NoSQL database only needs to give real implement to NoSqlGdbc interface and can give SqlGdbc fake (nil) implementation
// SQL database only needs to give real implement to SqlGdbc interface and can give NoSqlGdbc fake(nil) implementation
type Gdbc interface {
	SqlGdbc
	NoSqlGdbc
}

// Transactioner is the transaction interface for database handler
// It should only be applicable to SQL database
type Transactioner interface {
	// Rollback a transaction
	Rollback() error
	// Commit a transaction
	Commit() error
	// TxEnd commits a transaction if no errors, otherwise rollback
	// txFunc is the operations wrapped in a transaction
	TxEnd(txFunc func() error) error
	// TxBegin gets *sql.DB from receiver and return a SqlGdbc, which has a *sql.Tx
	TxBegin() (SqlGdbc, error)
}
