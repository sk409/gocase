package gocase

func UpperCamelCase(b []byte, golike bool) []byte {
	b = firstTransformer.upperCamelCase(b)
	b = snakeTransformer.upperCamelCase(b)
	if golike {
		b = golikeTransformer.upperCamelCase(b)
	}
	return b
}
