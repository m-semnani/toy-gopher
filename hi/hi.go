package hi

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// HelloWorld :)
func HelloWorld(name string) (string, error) {
	if name == "" {
		return "", errors.New("name can not be empty")
	}

	return fmt.Sprintf(getRandomMessage(), name), nil
}

func getRandomMessage() string {
	randMsg := []string{
		"Hi %v, welcome :)",
		"Hello there, nice to meet you %v :)",
		"Hey %v! Good morning.",
	}

	return randMsg[rand.Intn(len(randMsg))]
}
