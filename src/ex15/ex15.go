var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	b := make([]byte, 100)
	copy(b, v[:100])
	justString = string(b)
}

func main() {
	someFunc()
}