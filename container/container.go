package container

import "github.com/jfeng45/servicetmpl/usecase"

type Container interface {
	// InitApp loads the application configurations and logger
	 InitApp( filename string ) error

	// RetrieveRegistration retrieves RegistrationUseCaseInterface from factory map. If it is not in map yet,
	// it created one and put it into map.
	// RetrieveRegistration is a singleton
	 RetrieveRegistration() (usecase.RegistrationUseCaseInterface, error)

	// RetrieveListUser retrieves ListUserUseCaseInterface from factory map. If it is not in map yet, it created one
	// and put it into map.
	// RetrieveListUser is a singleton
	 RetrieveListUser() (usecase.ListUserUseCaseInterface, error)
}
