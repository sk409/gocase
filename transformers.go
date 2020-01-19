package gocase

import (
	"regexp"
	"strings"
)

type first struct {
}

func (f *first) lowerCamelCase(b []byte) []byte {
	if len(b) == 0 {
		return []byte{}
	}
	str := string(b)
	p := []string{"ID", "URL"}
	for _, v := range p {
		if strings.HasPrefix(str, v) {
			l := strings.ToLower(v)
			if len(v) == len(str) {
				return []byte(l)
			}
			return []byte(l + str[len(v):])
		}
	}
	l := strings.ToLower(string(str[0]))
	if len(str) == 1 {
		return []byte(l)
	}
	return []byte(l + str[1:])
}

func (f *first) upperCamelCase(b []byte) []byte {
	if len(b) == 0 {
		return []byte{}
	}
	str := string(b)
	u := strings.ToUpper(string(str[0]))
	if len(str) == 1 {
		return []byte(u)
	}
	return []byte(u + str[1:])
}

type golike struct {
}

func (g *golike) transform(b []byte) []byte {
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

func (u *upperCamel) lowerCamelCase(b []byte) []byte {
	runes := []rune(string(b))
	return []byte(string(runes))
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
