package greetings

import "errors"

var greetingMap = map[string]string{"english": "hello", "french": "bonjour", "spanish": "hola", "german": "guten tag", "hawaiian": "aloha e"}

func Greet(name string, language string) (string, error) {
	greeting, success := greetingMap[language]
	if !success {
		return "", errors.New("unknown language")
	}

	return greeting + " " + name + "!", nil
}
