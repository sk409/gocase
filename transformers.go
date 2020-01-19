package gocase

import (
	"regexp"
	"strings"
)

type first struct {
}

func (f *first) upperCamelCase(b []byte) []byte {
	if len(b) == 0 {
		return []byte{}
	}
	runes := []rune(string(b))
	runes[0] = rune(strings.ToUpper(string(b[0]))[0])
	return []byte(string(runes))
}

type golike struct {
}

func (g *golike) upperCamelCase(b []byte) []byte {
	str := string(b)
	str = strings.ReplaceAll(str, "Id", "ID")
	str = strings.ReplaceAll(str, "Url", "URL")
	return []byte(str)
}

type snake struct {
}

func (s *snake) upperCamelCase(b []byte) []byte {
	runes := []rune(string(b))
	regex := regexp.MustCompile("_([a-z])")
	matches := regex.FindAllSubmatchIndex(b, -1)
	for _, match := range matches {
		runes[match[2]] = rune(strings.ToUpper(string(b[match[2]]))[0])
	}
	str := string(runes)
	str = strings.ReplaceAll(str, "_", "")
	return []byte(str)
}

type upperCamel struct {
}

func (u *upperCamel) snakeCase(b []byte) []byte {
	regex := regexp.MustCompile("(.)([A-Z])")
	snake := regex.ReplaceAll(b, []byte("${1}_${2}"))
	return []byte(strings.ToLower(string(snake)))
}

var (
	firstTransformer      = first{}
	golikeTransformer     = golike{}
	snakeTransformer      = snake{}
	upperCamelTransformer = upperCamel{}
)
