// Package helpers define helper functions globally
package helpers

import (
	"os"
	"path/filepath"
	"runtime"
)

// GetAppEnv return application running mode
func GetAppEnv() string {
	switch os.Getenv("FIBER_ENV") {
	case "production":
		return "production"
	case "staging":
		return "staging"
	case "test":
		return "test"
	default:
		return "development"
	}
}

// IsProduction return true or false for production env
func IsProduction() bool {
	return GetAppEnv() == "production"
}

func AppRoot() string {
	_, callerFile, _, _ := runtime.Caller(0)
	executablePath := filepath.Dir(callerFile)
	appRootPath := filepath.Join(executablePath, "/../..")

	return appRootPath
}
