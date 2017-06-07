package main

import (
	"io/ioutil"
	"log"
	"strings"
	"os"
	"bufio"
)

func main() {
	f1 := "/Users/zhangguijiang/tmpDoc/logs/newgid"
	f2 := "/Users/zhangguijiang/tmpDoc/logs/gidIdfaFile"
	m := readFileToMap(f1)
	misList, multiList, nullList, matchList := diff(f2, m)
	log.Println(len(misList), len(multiList), len(nullList), len(matchList))
	writeFile("/Users/zhangguijiang/tmpDoc/logs/idfa_misList", misList)
	writeFile("/Users/zhangguijiang/tmpDoc/logs/idfa_multiList", multiList)
	writeFile("/Users/zhangguijiang/tmpDoc/logs/idfa_nullList", nullList)
	writeFile("/Users/zhangguijiang/tmpDoc/logs/idfa_matchList", matchList)
}

func readFileToMap(file string) map[string]string {
	bs, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal("read file err")
	}
	list := strings.FieldsFunc(string(bs), func(r rune) bool {
		return r == rune('\n')
	})
	m := make(map[string]string)
	for _, l := range list {
		ss := strings.Split(l, " ")
		if len(ss) != 2 {
			log.Println("readFileToMap num err line", ss)
			continue
		}
		m[string(ss[0])] = string(ss[1])
	}
	return m
}

func diff(file string, m map[string]string) ([]string, []string, []string, []string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal("open file err", err)
	}
	defer f.Close()
	// 不匹配的行
	misMatchList := make([]string, 0)
	// 多个gid匹配了其中一个的行
	multiList := make([]string, 0)
	// 新的文件中没有的行
	nullList := make([]string, 0)
	// 匹配的行
	matchList := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		s := scanner.Text()
		sl := strings.Fields(s)
		if len(sl) != 2 {
			//log.Println("diff num err line", sl)
			continue
		}
		gids := string(sl[0])
		idfa := string(sl[1])
		if gid, ok := m[idfa]; !ok {
			nullList = append(nullList, s+"\n")
		} else {
			if strings.Contains(gids, gid) {
				if strings.Contains(gids, ",") {
					gidL := strings.Split(gids, ",")
					ns := s + " ## " + gid + " ## "
					for i, g := range gidL {
						if g == gid {
							ns = ns + string(byte(i))
							break
						}
					}
					multiList = append(multiList, ns+"\n")
				} else {
					matchList = append(matchList, s+" ## "+gid+"\n")
				}
			} else {
				ns := s + " ## " + gid
				misMatchList = append(misMatchList, ns+"\n")
			}
		}
	}
	return misMatchList, multiList, nullList, matchList
}

func writeFile(filePath string, list []string) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal("create file err")
	}
	buf := bufio.NewWriter(file)
	for _, s := range list {
		buf.WriteString(s)
	}
	if err = buf.Flush(); err != nil {
		panic(err)
	}
}
