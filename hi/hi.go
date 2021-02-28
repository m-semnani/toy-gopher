package hi

import (
	"errors"
)

// HelloWorld :)
func HelloWorld(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name can not be empty!")
	}

	return "Hello " + name, nil
}
