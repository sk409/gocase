package gocase

func LowerCamelCase(b []byte, golike bool) []byte {
	if !golike {
		g := golikeTransformer{}
		b = g.lowerCamelCase(b)
	}
	u := upperCamelTransformer{}
	b = u.lowerCamelCase(b)
	return b
}

func SnakeCase(b []byte) []byte {
	u := upperCamelTransformer{}
	b = u.snakeCase(b)
	return b
}

func UpperCamelCase(b []byte, golike bool) []byte {
	// b = firstTransformer.upperCamelCase(b)
	// b = snakeTransformer.upperCamelCase(b)
	l := lowerCamelTransformer{}
	b = l.upperCamelCase(b)
	if golike {
		g := golikeTransformer{}
		b = g.upperCamelCase(b)
	}
	return b
}
