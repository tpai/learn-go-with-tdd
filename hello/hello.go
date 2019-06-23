package main

const helloPrefix = "Hello "
const chineseHelloPrefix = "你好 "
const japaneseHelloPrefix = "こんにちは "

const chinese = "Chinese"
const japanese = "Japanese"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (greeting string) {
	switch language {
	case chinese:
		greeting = chineseHelloPrefix
	case japanese:
		greeting = japaneseHelloPrefix
	default:
		greeting = helloPrefix
	}
	return
}
