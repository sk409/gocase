package gocase

func lowerCamelCase(b []byte, golike bool) []byte {
	b = upperCamelTransformer.lowerCamelCase(b)
	if golike {
		b = golikeTransformer.transform(b)
	}
	return b
}

func SnakeCase(b []byte) []byte {
	b = upperCamelTransformer.snakeCase(b)
	return b
}

func UpperCamelCase(b []byte, golike bool) []byte {
	b = snakeTransformer.upperCamelCase(b)
	if golike {
		b = golikeTransformer.transform(b)
	}
	return b
}
