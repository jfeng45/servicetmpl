package container

// use case code. Need to map to the use case code in the configuration yaml file.
const (
	REGISTRATION   string ="registration"
	LIST_USER string ="listUser"

)
type Container interface {
	// InitApp loads the application configurations, logger and concrete types for use cases interfaces
	InitApp( filename string ) error

	// GetInstance retrieves corresponding type base on code from factory map. It returns err if the type doesn't exist
	// GetInstance is a singleton
	// It is the out-facing interface used by business function
	GetInstance(code string) ( interface{}, error)

	// This should only be used by container and it's sub-package
	// Get instance by code from container
	Get(code string) (interface{}, bool)

	// This should only be used by container and it's sub-package
	// Put value into container with code as the key
	Put(code string, value interface{})

}
