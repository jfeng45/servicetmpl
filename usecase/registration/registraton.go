// Package registration represents the concrete implementation of RegistrationUseCaseInterface interface
// Because the same business function can be created to support both transaction and non-transaction, the business
// function need to be created to support both. So, the shared business function is created in a helper file and
// the wrapper of that function to support transaction and non-transaction are created in a file to
// implement use case
package registration

import (
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/pkg/errors"
)
// RegistrationUseCase implements RegistrationUseCaseInterface. It has UserDataInterface, which can be used to
// access persistence layer
type RegistrationUseCase struct {
	UserDataInterface  dataservice.UserDataInterface
}

func (uuc *RegistrationUseCase) RegisterUser(user *model.User) (*model.User, error) {
	err :=user.Validate()
	if err != nil {
		return nil, errors.Wrap(err, "user validation failed")
	}
	isDup, err := uuc.isDuplicate(user.Name)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	if isDup {
		return nil, errors.New("duplicate user for " + user.Name)
	}
	resultUser, err :=uuc.UserDataInterface.Insert(user)

	if err != nil {
		return nil, errors.Wrap(err, "")
	}
	return resultUser, nil
}

func (uuc *RegistrationUseCase) ModifyUser(user *model.User) error {
	return modifyUser(uuc.UserDataInterface, user)
}

func (uuc *RegistrationUseCase) isDuplicate(name string) (bool, error) {
	user, err :=uuc.UserDataInterface.FindByName(name)
	//logger.Log.Debug("isDuplicate() user:", user)
	if err != nil {
		return false, errors.Wrap(err, "")
	}
	if user != nil {
		return true, nil
	}
	return false, nil
}

func (uuc *RegistrationUseCase) UnregisterUser(username string) error {
	return unregisterUser(uuc.UserDataInterface, username)
}

// The use case of ModifyAndUnregister without transaction
func (uuc *RegistrationUseCase) ModifyAndUnregister(user *model.User) error {
	return modifyAndUnregister(uuc.UserDataInterface, user)
}
// The use case of ModifyAndUnregister with transaction
func (uuc *RegistrationUseCase) ModifyAndUnregisterWithTx(user *model.User) error {
	newUdi, err := uuc.UserDataInterface.TxBegin()
	if err != nil {
		return errors.Wrap(err, "")
	}
	return newUdi.TxEnd( func () error {
		//wrap the business function inside the TxEnd function
		return modifyAndUnregister(newUdi, user)
	})
}
