// Package usecasefactory using factory method pattern to create concrete case case.
// Generally speaking, each use case needs a separate factory.
package usecasefactory

import (
	"github.com/jfeng45/servicetmpl/configs"
	"github.com/jfeng45/servicetmpl/container"
)

//To map "use case code" to "use case interface builder"
// Each use case has exactly one factory. For example, "registration" use case has "RegistrationFactory"
// Each factory has it's own file. For example, "RegistrationFactory" is in "registrationFactory.go"
var UseCaseFactoryBuilderMap = map[string]UseCaseFbInterface {
	container.REGISTRATION: &RegistrationFactory{},
	container.LIST_USER: &ListUserFactory{},
	container.LIST_COURSE: &ListCourseFactory{},
}

// UseCaseInterface serve as a marker to indicate the return type for Build method
type UseCaseInterface interface{}

// The builder interface for factory method pattern
// Every factory needs to implement build method
type UseCaseFbInterface interface {
	Build(c container.Container, appConfig *configs.AppConfig, key string ) (UseCaseInterface, error)
}

//GetDataStoreFb is accessors for factoryBuilderMap
func GetUseCaseFb(key string) UseCaseFbInterface {
	return UseCaseFactoryBuilderMap[key]
}




