// Package registration represents the concrete implementation of ListUserUseCaseInterface interface
package listuser

import (
	"github.com/jfeng45/servicetmpl/container/logger"
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/pkg/errors"
	"strconv"
)

// ListUserUseCase implements ListUseCaseInterface.
type ListUserUseCase struct {
	// UserDataInterface, which is a interface to underline database connection and can be used to access
	// persistence layer
	UserDataInterface  dataservice.UserDataInterface
	// CacheDataInterface, which is a interface to outside gRPC cache service and can be used to access gRPC service
	CacheDataInterface dataservice.CacheDataInterface
}

func (luc *ListUserUseCase) ListUser() ([]model.User, error) {
	return luc.UserDataInterface.FindAll()
}
func (luc *ListUserUseCase)Find(id int) (*model.User,error) {
	users, err := luc.getFromCache(strconv.Itoa(id))
	if err != nil {
		//not found in cache and continue
		logger.Log.Errorf("get from cache error:", err)
		//return nil, errors.Wrap(err, "")
	}
	if users != nil {
		//here should return the results from cache, however, right now the cache doesn't store user info,
		//so, just call find(id). This is not real code. Please replace it with real code
		return luc.UserDataInterface.Find(id)
	}
	return luc.UserDataInterface.Find(id)
}

// GetFromCache is a fake function to just show a call to outside service, the call to outside is working,
// but the results returned is not right
func (luc *ListUserUseCase) getFromCache(key string) ([]model.User, error) {
	value, err := luc.CacheDataInterface.Get(key)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	logger.Log.Info("value from get cache: ", value)
	return nil, nil
}

