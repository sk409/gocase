package gocase

func LowerCamelCase(b []byte, golike bool) []byte {
	b = firstTransformer.lowerCamelCase(b)
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
	b = firstTransformer.upperCamelCase(b)
	b = snakeTransformer.upperCamelCase(b)
	if golike {
		b = golikeTransformer.transform(b)
	}
	return b
}
