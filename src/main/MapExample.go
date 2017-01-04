package main

import "fmt"

func main() {
	//初始化map，key类型为string，value类型为string
	//等同于 m := map[string]string{}
	m := make(map[string]string)

	//向map中添加key-value
	m["1"] = "v1"
	m["2"] = "v2"
	m["3"] = "v3"
	fmt.Println(m)
	fmt.Println(len(m))

	//获取key为 "1" 的value
	s1 := m["1"]
	fmt.Println(s1)

	//获取key为 "4" 的value
	//map中实际不存在这个key，获取到的value为?
	s2 := m["4"]
	fmt.Println(s2)
	fmt.Println(s2 == "")
	fmt.Println(len(s2))

	//获取map中value时，根据返回的bool值来确定map中是否包含该key
	s1, ok := m["1"]
	fmt.Println("value: ", s1, "ok: ", ok)

	s2, ok = m["4"]
	fmt.Println("value: ", s2, "ok: ", ok)

	delete(m, "3")
	fmt.Println(m)

	//通过map实现set
	set := map[string]bool{
		"k1": true,
		"k2": true,
		"k3": true,
		//"k1": true,
	}
	fmt.Println(set)
}
