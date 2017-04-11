package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "aabcddgfdgjklrehgkuuerhuihgkehrkugahrekghreitughrtiuhngvkarjf" +
		"kghnketjhgktureghnjkerhtgnketrhgnktjrehgkrethgk" +
		"guioerhgitrhgkretkgjrtkhgeuirhjgvkeruhjnkghnretkg" +
		"ghjiokeruhgirethgkergkljfdklgeiurtghirt" +
		"ghieuhgiurtoignjsdflknavgkdfhngkjshnkljfbgkihtguirt" +
		"ghiuehgiurthgkjfsdanmlskjreuightreuikajhnlgjnkdjfgheiurhgie" +
		"gherhguiehjiugkjadkghnkjdfhgkjhguiregdjnfskaghkerlhrkg"
	sep := "bcddgfdgjklrehgkuuerhuihgkehrkugahrekghreitughrtiuhngvkarjf" +
		"kghnketjhgktureghnjkerhtgnketrhgnktjrehgkrethgk" +
		"guioerhgitrhgkretkgjrtkhgeuirhjgvkeruhjnkghnretkg" +
		"ghjiokeruhgirethgkergkljfdklgeiurtghirt" +
		"ghieuhgiurtoignjsdflknavgkdfhngkjshnkljfbgkihtguirt" +
		"ghiuehgiurthgkjfsdanmlskjreuightreuikajhnlgjnkdjfgheiurhgie" +
		"gherhguiehjiugkjadkghnkjdfhgkjhguiregdjnfskaghkerlhrkg"
	i := strings.Index(s, sep)
	fmt.Println(i)
}
