// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package listuser

import (
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/pkg/errors"
	"strconv"
)

// RegistrationUseCase implements RegistrationUseCaseInterface.

type ListUserUseCase struct {
	// UserDataInterface, which is a interface to underline database connection and can be used to access
	// persistence layer
	UserDataInterface  dataservice.UserDataInterface
	// CacheDataInterface, which is a interface to outside gRPC cache service and can be used to access gRPC service
	CacheDataInterface dataservice.CacheDataInterface
}

func (uuc *ListUserUseCase) ListUser() ([]model.User, error) {
	return uuc.UserDataInterface.FindAll()
}
func (uuc *ListUserUseCase)Find(id int) (*model.User,error) {
	users, err := uuc.getFromCache(strconv.Itoa(id))
	if err != nil {
		//not find and continue
		logger.Log.Errorf("get from cache error:", err)
		//return nil, errors.Wrap(err, "")
	}
	if users != nil {
		//here should return the results from cache, however, right now the cache doesn't store user info,
		//so, just call find(id). This is not real code. Please replace it with real code
		return uuc.UserDataInterface.Find(id)
	}
	return uuc.UserDataInterface.Find(id)
}

// getFromCache is a fake function to just show a call to outside service, the call to outside is working, but the results returned is not right
func (uuc *ListUserUseCase) getFromCache(key string) ([]model.User, error) {
	value, err := uuc.CacheDataInterface.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Info("value from get cache: ", value)
	return nil, nil
}

