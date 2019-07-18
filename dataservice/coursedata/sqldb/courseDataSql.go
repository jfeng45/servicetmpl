package sqldb

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/jfeng45/servicetmpl/tools/gdbc"
	"github.com/pkg/errors"
)

const (
	QUERY_COURSE = "SELECT * FROM course"
)
// CourseDataSql is the SQL implementation of CourseDatainterface
type CourseDataSql struct {
	DB gdbc.SqlGdbc
}

func (cds *CourseDataSql) SetDB(gdbc gdbc.Gdbc) {
	cds.DB = gdbc
}

func (cds *CourseDataSql) FindAll() ([]model.Course, error) {

	rows, err := cds.DB.Query(QUERY_COURSE)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	defer rows.Close()
	courses := []model.Course{}

	//var ds string
	for rows.Next() {
		course, err :=rowsToCourse(rows)
		if err != nil {
			return courses, errors.Wrap(err,"")
		}
		courses = append(courses, *course)

	}
	//need to check error for rows.Next()
	if err =rows.Err(); err!= nil {
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Debug("find course list:", courses)
	return courses, nil
}

func rowsToCourse(rows *sql.Rows) (*model.Course, error) {
	course := &model.Course{}
	err := rows.Scan(&course.Id, &course.Name)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	logger.Log.Debug("rows to Course:", course)
	return course, nil
}





