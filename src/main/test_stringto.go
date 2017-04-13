package main

import (
	"unicode"
	"strings"
	"fmt"
)

func main() {
	lower := "abcçdefgğhıijklmnoöprsştuüvyz中"
	upper := "ABCÇDEFGĞHIİJKLMNOÖPRSŞTUÜVYZ国"
	u := strings.ToUpperSpecial(unicode.TurkishCase, lower)
	l := strings.ToLowerSpecial(unicode.TurkishCase, upper)
	fmt.Println(u)
	fmt.Println(strings.ToUpper(lower))
	fmt.Println(l)
	fmt.Println(strings.ToLower(upper))
}