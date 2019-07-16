package registration

import (
	"github.com/jfeng45/servicetmpl/dataservice"
	"github.com/jfeng45/servicetmpl/model"
	"github.com/pkg/errors"
	"strconv"
)

func modifyUser(udi dataservice.UserDataInterface, user *model.User) error {
	//loggera.Log.Debug("modifyUser")
	err :=user.ValidatePersist()
	if err != nil {
		return errors.Wrap(err, "user validation failed")
	}
	rowsAffected, err := udi.Update(user)
	if err!= nil {
		return errors.Wrap(err, "")
	}
	if rowsAffected != 1 {
		return errors.New("Modify user failed. rows affected is " + strconv.Itoa(int(rowsAffected)))
	}
	return nil
}

func unregisterUser(udi dataservice.UserDataInterface, username string) error {
	affected, err := udi.Remove(username)
	if err != nil {
		return errors.Wrap(err, "")
	}
	if affected == 0 {
		errStr := "UnregisterUser failed. No such user " + username
		return errors.New(errStr)
	}

	if affected != 1 {
		errStr :="UnregisterUser failed. Number of users unregistered are  " + strconv.Itoa(int(affected))
		return errors.New(errStr)
	}
	return nil
}

// The business function will be wrapped inside a transaction and without a transaction
// It needs to be written in a way that every error will be returned so it can be catched by TxEnd() function,
// which will handle commit and rollback
func modifyAndUnregister(uuc *RegistrationUseCase, user *model.User) error {
	//loggera.Log.Debug("modifyAndUnregister")
	udi := uuc.UserDataInterface
	err := modifyUser(udi, user)
	if err != nil {
		return errors.Wrap(err, "")
	}
	err = unregisterUser(udi, user.Name)
	if err != nil {
		return errors.Wrap(err, "")
	}
	return nil
}