package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if len(name) == 0 {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Greate to see you, %v",
		"Gratias %v. Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
