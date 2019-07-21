package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/config"
	"github.com/jfeng45/servicetmpl/container"
	"github.com/jfeng45/servicetmpl/usecase/listcourse"
	"github.com/pkg/errors"
)

type ListCourseFactory struct{}

func (lcf *ListCourseFactory) Build(c container.Container, appConfig *config.AppConfig, key string) (UseCaseInterface, error) {
	uc := appConfig.UseCase.ListCourse

	if container.LIST_COURSE != uc.Code {
		errMsg := container.LIST_COURSE + " in ListCourseFactory doesn't match key = " + key
		return nil, errors.New(errMsg)
	}

	cdi, err := buildCourseData(c, &uc.CourseDataConfig)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	lcuc := listcourse.ListCourseUseCase{CourseDataInterface: cdi}
	return &lcuc, nil
}
