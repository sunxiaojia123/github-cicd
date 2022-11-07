package main

import "fmt"

func main() {
	saying := Cat()
	fmt.Println(saying)
}

// cat 方法
func Cat() string {
	return "miao~~~~~~~"
}
