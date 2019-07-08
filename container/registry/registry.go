// package registry represents application configuration registry, so the configuration is loaded only once
package registry

import (
	"github.com/jfeng45/servicetmpl/tools/logger"
)

// GetFromRegistry get configuration from registry based on key
func GetFromRegistry(factoryMap map[string]interface{}, key string) (interface{}, bool) {

	logger.Log.Debug("getFromRegistry: key=", key)
	luc1, found := factoryMap[key]
	return luc1, found

}

