package main

import (
	"fmt"
	"strings"
)

func main() {
	broken := "F#ck sl#t"
	replacer := strings.NewReplacer("#", "u")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

}

// Работа с исправлением символов используя метод strings.

