package main

import (
	"io/ioutil"
	"log"
	"strings"
	"unicode"
	"fmt"
)

func main() {
	s := "VACATION_SPECIAL_MAP.put(\"###\", \"$$$\");"
	vacation := readVacation("/Users/zhangguijiang/tmpDoc/logs/vacation.txt")
	des := readDes("/Users/zhangguijiang/tmpDoc/logs/destination.txt")
	//log.Println(len(vacation), len(des))
	//ndes := des[:1]
	//.Println(ndes)
	for _,d:=range des{
		for _,v:=range vacation{
			//fmt.Println(v)
			if v == "" {
				continue
			}
			if strings.EqualFold(d, v) {
				//fmt.Println(d, v)
				break
			}
			if strings.Contains(d, v) || strings.Contains(v, d) {
				//fmt.Println(d, v)
				st := strings.Replace(s, "###", d, -1)
				st = strings.Replace(st, "$$$", v, -1)
				fmt.Println(st)
				break
			}
		}
	}
}

func readVacation(filePath string) []string {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("read vacation err")
	}
	list := strings.FieldsFunc(string(bs), func(c rune) bool {
		return unicode.IsSpace(c) || string(c) == ":"
	})
	m := make(map[string]int, len(list))
	for _, k := range list{
		m[k] = 1
	}
	l := make([]string, 0)
	for k := range m{
		//log.Println(k)
		l = append(l, k)
	}
	return l
}

func readDes(filePath string) []string {
	bs, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("read vacation err")
	}
	list := strings.FieldsFunc(string(bs), func(c rune) bool {
		return unicode.IsSpace(c)
	})
	return list
}