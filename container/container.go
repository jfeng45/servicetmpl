// package container use dependency injection to create concrete type and wire the whole application together
package container

// use case code. Need to map to the use case code (UseCaseConfig) in the configuration yaml file.
// Client app use those to retrieve use case from the container
const (
	REGISTRATION   string ="registration"
	LIST_USER string ="listUser"
	LIST_COURSE string ="listCourse"

)
type Container interface {
	// InitApp loads the application configurations from a file, initialize logger and concrete types for the application
	InitApp( filename string ) error

	// GetInstance retrieves corresponding type base on code from factory map. It returns err if the type doesn't exist
	// GetInstance is a singleton
	// It is the out-facing interface used by business function
	//GetInstance(code string) ( interface{}, error)

	BuildUseCase(code string) ( interface{}, error)

	// This should only be used by container and it's sub-package
	// Get instance by code from container
	Get(code string) (interface{}, bool)

	// This should only be used by container and it's sub-package
	// Put value into container with code as the key
	Put(code string, value interface{})

}
