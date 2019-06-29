package main

import (
	cw "./concurrency"
	"fmt"
)

func main() {
	results := cw.CheckWebsites(cw.CheckWebsite, []string{
		"https://google.com",
		"https://wtfismyip.com",
		"wazup://whats.up",
	})
	fmt.Printf("%v", results)
}
