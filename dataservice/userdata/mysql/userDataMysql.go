// Package mysql represents the mySQL implementation of the user data persistence layer
package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/dataservice/gdbc"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/jfeng45/servicetmpl/tools"
	"github.com/jfeng45/servicetmpl/tools/logger"
	"github.com/pkg/errors"
	"time"
)

// DBTxStore is the MySQL implementation of Gdbc interface
type DBTxStore struct {
	DB gdbc.Gdbc
}

func (userData *DBTxStore) Remove(username string) (int64, error) {

	stmt, err := userData.DB.Prepare("delete from userinfo where username=?")
	if err!=nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()

	res, err := stmt.Exec(username)
	if err!=nil {
		return 0, errors.Wrap(err, "")
	}
	rowsAffected, err := res.RowsAffected()
	if err!=nil {
		return 0, errors.Wrap(err, "")
	}

	logger.Log.Debug("remove:row affected ", rowsAffected)
	return rowsAffected, nil
}

func (userData *DBTxStore) Find(id int) (*model.User, error) {
	rows, err := userData.DB.Query("SELECT * FROM userinfo where uid =?", id)
	if err!=nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	return retrieveUser(rows)
	//return user, nil
}
func retrieveUser(rows *sql.Rows) (*model.User, error) {
	if rows.Next() {
		return rowsToUser(rows)
	}
	return nil, nil
}
func rowsToUser(rows *sql.Rows) (*model.User, error) {
	var ds string
	user := &model.User{}
	err := rows.Scan(&user.Id, &user.Name, &user.Department, &ds)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	created, err := time.Parse(tools.FORMAT_ISO8601_DATE, ds)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	user.Created = created

	logger.Log.Debug("rows to User:", user)
	return user, nil
}
func (userData *DBTxStore) FindByName(name string) (*model.User, error) {
	rows, err := userData.DB.Query("SELECT * FROM userinfo where username =?", name)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	return retrieveUser(rows)
}

func (userData *DBTxStore) FindAll() ([]model.User, error) {

	rows, err := userData.DB.Query("SELECT * FROM userinfo ")
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	users := []model.User{}

	//var ds string
	for rows.Next() {
		user, err :=rowsToUser(rows)
		if err != nil {
			return users, errors.Wrap(err,"")
		}
		users = append(users, *user)

	}
	//need to check error for rows.Next()
	if err =rows.Err(); err!= nil {
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("find user list:", users)
	return users, nil
}

func (userData *DBTxStore) Update(user *model.User) (int64, error) {

	stmt, err := userData.DB.Prepare("update userinfo set username=?, department=?, created=? where uid=?")

	if err!=nil {
		return 0, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Name, user.Department, user.Created, user.Id)
	if err!=nil {
		return 0, errors.Wrap(err, "")
	}
	rowsAffected, err := res.RowsAffected()

	if err!=nil {
		return 0, errors.Wrap(err, "")
	}
	logger.Log.Debug("update: rows affected: ", rowsAffected)

	return rowsAffected, nil
}

func (userData *DBTxStore) Insert(user *model.User) (*model.User, error) {

	stmt, err := userData.DB.Prepare("INSERT userinfo SET username=?,department=?,created=?")
	if err!=nil {
		return nil, errors.Wrap(err, "")
	}
	defer stmt.Close()
	res, err := stmt.Exec(user.Name, user.Department, user.Created)
	if err!=nil {
		return nil, errors.Wrap(err, "")
	}
	id, err := res.LastInsertId()
	if err!=nil {
		return nil, errors.Wrap(err, "")
	}
	user.Id = int(id)
	logger.Log.Debug("user inserted:", user)
	return user, nil
}

func (dataStore *DBTxStore)TxEnd( txFunc func() error) error {
	return dataStore.DB.TxEnd(txFunc)
}

func (dataStore *DBTxStore)TxBegin() (dataservice.UserDataInterface, error) {

	gdbc, error :=dataStore.DB.TxBegin()
	dbts := DBTxStore{gdbc}
	return &dbts, error
}