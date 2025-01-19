package greetings

// For more on naming your module with a module path, see Managing dependencies
// https://go.dev/doc/modules/managing-dependencies#naming_module
import (
	"errors"
	"fmt"
	"math/rand"
)

// it only works inside of package
func randomMessage() string {
	messages := []string{
		"Hi, %v. Welcome to my module",
		"Nice to meet you, %v",
		"Whats up dude %v",
		"Hello world! Outstanding %v",
	}

	return messages[rand.Intn(len(messages))]
}

// exported name
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome to my module", name)
	return message, nil
}

func RandomHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomMessage(), name)
	return message, nil
}
