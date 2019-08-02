package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	source := strings.Join(os.Args[1:2], "")
	dupSubject := strings.Join(os.Args[2:3], "")

	if source == "" || dupSubject == "" {
		panic("We need a source string and a char to duplicate")
	}

	duplicate := dupSubject + dupSubject
	result := strings.Replace(source, dupSubject, duplicate, -1)

	fmt.Println(result)
}
