package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Nombre vacío")
	}

	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!\n",
		"Great to see you, %v!\n",
		"Hail, %v! Well met!\n",
		"Hola %v, ¿Como estas?\n",
		"Saludos %v! ¿Cómo te va?\n",
		"Bienvenido %v! ¿Qué tal?\n",
	}

	return formats[rand.Intn(len(formats))]
}
