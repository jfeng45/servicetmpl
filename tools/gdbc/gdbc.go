// package gdbc is created to represents low level database interfaces in order to have a unified way to handle data issues.
// For example, database transactions, interface for SQL and NoSQL database.
package gdbc

// gdbc is a unified way to handle database connections. As long as a data service layer type implements this interaface,
// it can work with both SQL and NOSQL underline database.
// NoSQL database only needs to implement NoSqlGdbc and can give SqlGdbc nil implementation
// SQL database only needs to implement SqlGdbcNoSqlGdbc and can give NoSqlGdbc nil implementation
type gdbc interface {
	SqlGdbc
	NoSqlGdbc
}
