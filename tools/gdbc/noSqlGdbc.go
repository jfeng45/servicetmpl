package gdbc

import (
	"context"
	//"github.com/flimzy/kivik"
	"github.com/go-kivik/kivik"
)

// NoSqlGdbc (NoSQL Go database connection) is a wrapper for NOSql database handler
// Currently, it is just a illustration of an idea instead of a workable implementation
// The real implementation will remove kivik related type and need to do conversion between generic types and
// individual types
type NoSqlGdbc interface {
	// The method name was Query(), but since it conflict with the name with Query() in SqlGdbc,
	// so changed to a different name
	QueryNoSql(ctx context.Context,ddoc string, view string) (*kivik.Rows, error)
	Put(ctx context.Context, docID string, doc interface{}, options ...kivik.Options) (rev string, err error)
	Get(ctx context.Context, docID string, options ...kivik.Options) (*kivik.Row, error)
	Find(ctx context.Context, query interface{}) (*kivik.Rows, error)
	AllDocs(ctx context.Context, options ...kivik.Options ) (*kivik.Rows, error)
}

// NoSqlDB is the concrete implementation of GDBC by using *sql.DB
type NoSqlDB struct {
	DB *kivik.DB
}

func (nsdb *NoSqlDB )QueryNoSql(ctx context.Context, ddoc string, view string)  (*kivik.Rows, error){
	return nsdb.DB.Query(ctx ,ddoc , view)
}

func (nsdb *NoSqlDB )Put(ctx context.Context, docID string, doc interface{}, options ...kivik.Options) (rev string, err error) {
	return nsdb.DB.Put(ctx, docID, doc, options ...)
}

func (nsdb *NoSqlDB )Get(ctx context.Context, docID string, options ...kivik.Options) (*kivik.Row, error) {
	return nsdb.DB.Get(ctx, docID, options ...)
}

func (nsdb *NoSqlDB )Find(ctx context.Context, query interface{}) (*kivik.Rows, error) {
	return nsdb.DB.Find(ctx, query)
}

func (nsdb *NoSqlDB )AllDocs(ctx context.Context, options ...kivik.Options ) (*kivik.Rows, error) {
	return nsdb.DB.AllDocs(ctx, options ... )
}



