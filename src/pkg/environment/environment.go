// Package environment defines helpers accessing environment values
package environment

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

// Env represents environmental variable instance
type Env struct{}

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
