// Package environment defines helpers accessing environment values
package environment

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Env represents environmental variable instance
type Env struct {
	envCache map[string]string
}

// New creates a new instance of Env and returns an error if any occurs
func New() (*Env, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	return &Env{}, nil
}

// NewLoadFromFile lets you load Env object from a file
func NewLoadFromFile(fileName string) (*Env, error) {
	err := godotenv.Load(fileName)
	if err != nil {
		return nil, err
	}
	return &Env{}, nil
}

// Get retrieves the string value of an environmental variable
func (e *Env) Get(key string) string {
	return os.Getenv(key)
}

// UseMock is helper that returns true or false if the environment should use mocks when hitting 3rd party partners
func (e *Env) UseMock() bool {
	v := e.Get("APP_MOCK")
	if len(v) == 0 {
		return false
	}

	if strings.EqualFold(v, "true") {
		return true
	}

	return false
}

// HelperForMocking [do not use in logic] designed for mocking and in test suite
func (e *Env) HelperForMocking(cache map[string]string) {
	e.envCache = cache
}

// IsSandbox is helper that returns true or false if the environment is in sandbox/testing mode
func (e *Env) IsSandbox() bool {
	v := e.Get("IS_SANDBOX_MODE")
	return strings.EqualFold(v, "true")
}
