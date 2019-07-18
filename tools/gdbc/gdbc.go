// package gdbc is created to represents low level database interfaces in order to have an unified way to
// access database handler.
// It is created to make it easier to handle certain database operations like transactions, database factory.
// It is ony a POC, not a mature solution
package gdbc

// gdbc is an unified way to handle database connections. As long as a data service type implements this interface,
// it can work with both SQL and NOSQL underline database.
// NoSQL database only needs to give real implement to NoSqlGdbc interface and can give SqlGdbc fake (nil) implementation
// SQL database only needs to give real implement to SqlGdbc interface and can give NoSqlGdbc fake(nil) implementation
type Gdbc interface {
	SqlGdbc
	NoSqlGdbc
}
