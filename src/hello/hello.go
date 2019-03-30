package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

// (prefix string) 命名返回值
// 将在函数中创建一个名为 prefix 的变量
// 它将被分配“零”值。这取决于类型，例如，int 是 0，string 是 ""
// 只需调用 return 而不是 return prefix 即可返回所设置的值
// 这将在godoc中显示，以便让代码更清晰
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = helloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
