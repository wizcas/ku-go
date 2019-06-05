package config

import (
	"fmt"
	"os"
	"strconv"
)

// EnvVar defines the environment varaible to retrieve
type EnvVar struct {
	// Key is the name of environment variable
	Key string
	// Required determines whether to be panic if environment variable not found.
	// It uses the `DefaultValue` if the variable is not required.
	Required bool
	// NotEmpty determines whether to be panic if the variable value is an empty string.
	NotEmpty bool
	// Default Value is the value returned when the variable is not `Required`
	DefaultValue string
}

// GetString returns the variable as a string value
func (ev EnvVar) GetString() string {
	val, ok := os.LookupEnv(ev.Key)
	if !ok {
		if ev.Required {
			panic(ev.errNotFound())
		} else {
			val = ev.DefaultValue
		}
	}
	if ev.NotEmpty && len(val) == 0 {
		panic(ev.errEmpty())
	}
	return val
}

// GetInt returns the variable as a integer value.
// It panics when the value assigned cannot be converted to integer,
// e.g. an empty string, a string that contains non-digit characters, etc.
func (ev EnvVar) GetInt() int {
	strval := ev.GetString()
	if len(strval) == 0 {
		panic(ev.errInvalidValue("cannot convert empty string into an integer"))
	}
	if val, err := strconv.ParseInt(strval, 10, 32); err != nil {
		panic(ev.errInvalidValue(err.Error()))
	} else {
		return int(val)
	}
}

func (ev EnvVar) errNotFound() error {
	return fmt.Errorf("environment variable is required: %s", ev.Key)
}
func (ev EnvVar) errEmpty() error {
	return fmt.Errorf("environment variable '%s' must not be empty", ev.Key)
}
func (ev EnvVar) errInvalidValue(err string) error {
	return fmt.Errorf("Invalid value of environment variable '%s': %v", ev.Key, err)
}
