package gocase

import (
	"bytes"
	"regexp"
	"strings"
)

// type first struct {
// }

// func (f *first) lowerCamelCase(b []byte) []byte {
// 	if len(b) == 0 {
// 		return []byte{}
// 	}
// 	str := string(b)
// 	p := []string{"ID", "URL", "SHA"}
// 	for _, v := range p {
// 		if strings.HasPrefix(str, v) {
// 			l := strings.ToLower(v)
// 			if len(v) == len(str) {
// 				return []byte(l)
// 			}
// 			return []byte(l + str[len(v):])
// 		}
// 	}
// 	l := strings.ToLower(string(str[0]))
// 	if len(str) == 1 {
// 		return []byte(l)
// 	}
// 	return []byte(l + str[1:])
// }

// func (f *first) upperCamelCase(b []byte) []byte {
// 	if len(b) == 0 {
// 		return []byte{}
// 	}
// 	str := string(b)
// 	u := strings.ToUpper(string(str[0]))
// 	if len(str) == 1 {
// 		return []byte(u)
// 	}
// 	return []byte(u + str[1:])
// }

type golikeTransformer struct {
}

func (g *golikeTransformer) lowerCamelCase(b []byte) []byte {
	str := string(b)
	// log.Println(str)
	golikeWords := []string{"ID", "URL", "SHA"}
	for _, golikeWord := range golikeWords {
		pointer := 0
		for {
			substr := str[pointer:]
			index := strings.Index(substr, golikeWord)
			if index == -1 {
				break
			}
			// log.Println(index)
			head := pointer + index
			tail := head + len(golikeWord)
			// log.Println(head)
			// log.Println(tail)
			ok := false
			if head == 0 || len(str) == tail {
				ok = true
			} else {
				pre := str[head-1 : head]
				next := str[tail : tail+1]
				// log.Println(pre)
				// log.Println(next)
				if strings.ToLower(pre) == pre && strings.ToLower(next) == next {
					ok = true
				}
			}
			// log.Println(ok)
			if ok {
				str = str[:head] + strings.ToUpper(golikeWord[0:1]) + strings.ToLower(golikeWord[1:]) + str[tail:]
			}
			// log.Println(str)
			// log.Println("====")
			pointer = head + len(golikeWord)
		}
	}
	return []byte(str)
}

func (g *golikeTransformer) upperCamelCase(b []byte) []byte {
	str := string(b)
	// log.Println(str)
	golikeWords := []string{"Id", "Url", "Sha"}
	for _, golikeWord := range golikeWords {
		pointer := 0
		for {
			substr := str[pointer:]
			index := strings.Index(substr, golikeWord)
			if index == -1 {
				break
			}
			// log.Println(index)
			head := pointer + index
			tail := head + len(golikeWord)
			// log.Println(head)
			// log.Println(tail)
			ok := false
			if index == 0 || len(str) == tail {
				ok = true
			} else {
				pre := str[head-1 : head]
				next := str[tail : tail+1]
				// log.Println(pre)
				// log.Println(next)
				if strings.ToLower(pre) == pre && strings.ToUpper(next) == next {
					ok = true
				}
			}
			if ok {
				str = str[:head] + strings.ToUpper(golikeWord) + str[tail:]
			}
			pointer = head + len(golikeWord)
		}
	}
	return []byte(str)
}

type lowerCamelTransformer struct {
}

func (l *lowerCamelTransformer) upperCamelCase(b []byte) []byte {
	if len(b) <= 1 {
		return b
	}
	return append(bytes.ToUpper(b[0:1]), b[1:]...)
}

type snakeTransformer struct {
}

func (s *snakeTransformer) upperCamelCase(b []byte) []byte {
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

type upperCamelTransformer struct {
}

func (u *upperCamelTransformer) lowerCamelCase(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	str := string(b)
	p := []string{"ID", "URL", "SHA"}
	for _, v := range p {
		if strings.HasPrefix(str, v) {
			l := strings.ToLower(v)
			if len(v) == len(str) {
				return []byte(l)
			}
			return []byte(l + str[len(v):])
		}
	}

	return append(bytes.ToLower(b[0:1]), b[1:]...)
}

func (u *upperCamelTransformer) snakeCase(b []byte) []byte {
	regex := regexp.MustCompile("([^A-Z])([A-Z])")
	snake := regex.ReplaceAll(b, []byte("${1}_${2}"))
	return []byte(strings.ToLower(string(snake)))
}
