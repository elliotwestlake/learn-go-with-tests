package main

var acceptedLanguages = map[string]string{
	"Spanish": "Hola",
	"French":  "Bonjour",
	"English": "Hello",
	"German":  "Hallo",
}

type Params struct {
	Name     string
	Language string
}

func Hello(prm Params) string {
	if prm.Language == "" {
		prm.Language = "English"
	}

	if prm.Name == "" {
		prm.Name = "World"
	}

	return acceptedLanguages[prm.Language] + ", " + prm.Name + "!"
}
