package main

import (
	"fmt"
)

const (
	defaultHello = "Hello "
	defaultName  = "World"
)

var helloIn = map[string]string{
	"en-GB": "Hello ",
	"es-ES": "Hola ",
	"fr-FR": "Bonjour ",
}

func main() {
	fmt.Println(hello("Kevin", "en-GB"))
}

func hello(name, locale string) string {
	if name == "" && locale == "" {
		return "Hello World!"
	}
	if name == "" {
		return fmt.Sprintf("%s%s!", helloIn[locale], defaultName)
	}
	if locale == "" {
		return fmt.Sprintf("%s%s!", defaultHello, name)
	}
	return fmt.Sprintf("%s%s!", helloIn[locale], name)
}
