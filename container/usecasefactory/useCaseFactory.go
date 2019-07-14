package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
)
// builder map to use case code to use case interface builder
// each use case has exactly one factory. For example, "registration" use case has "RegistrationFactory"
// Each factory has it's own file. For example, "RegistrationFactory" as in "registrationFactory.go"
var UseCaseFactoryBuilderMap = map[string]UseCaseFbInterface {
	container.REGISTRATION: &RegistrationFactory{},
	container.LIST_USER: &ListUserFactory{},
}

// The builder interface for factory method pattern
// Every factory needs to implement build method
type UseCaseFbInterface interface {
	Build(c container.Container, appConfig *configs.AppConfig, key string ) (err error)
}




