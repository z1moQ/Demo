package main //程序的包名

/*
单行调用写法，多个包时不推荐
import "fmt"
import "time"
*/

import (
	"fmt"
	"time"
)

//main函数
func main() { //花括号"{"必须和函数名在同一行，否则会编译失败
	//golang中的表达式加不加";"都可以，最好不加
	fmt.Println("Helo World")
	time.Sleep(1 * time.Second)
}
