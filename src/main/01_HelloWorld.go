// 声明本文件的package名
package main

// 倒入文件中使用到的package
// 注意：如果没有使用导入的package，必须删除或注释掉，否则编译不通过
import (
	"fmt"
)

func main() {
	fmt.Println("Hello world")
}
