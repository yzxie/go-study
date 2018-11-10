package main

import (
	"fmt"
	"log"
	"structs/exportStruct"
)

/*
 * 嵌套结构体
 */
 type Person struct {
	name string
	age int
	Address //存在字段提升，即可通过address访问
}
type Address struct {
	city string
}

/*
 * 包含匿名字段结构体
 */
type Anno struct {
	string
	int
}

/*
 * 包级别变量
 */
 var maxAge int = 120

/*
 * 初始化方法
 */
func init() {
	if maxAge > 120 {
		log.Fatal("width is invalid")
	}
}

func main() {
	/*
	 * 导入的结构体
	 */
	// 结构体创建实例
	var e1 exportStruct.Employee // 声明
	// 赋值
	e1 = exportStruct.Employee{
		Name: "xieyizun",
		Email: "123@164.com",
		Age: 25,
	}
	e2 := exportStruct.Employee{"xieyizun2", "xx", 26}
	fmt.Println("e1, e2", e1, e2)
	// 访问结构体字段
	fmt.Println("e1 name", e1.Name)

	// 嵌套结构体
	var person Person
	person.name = "xyz"
	person.age = 27
	person.Address = Address{
		city: "guangzhou",
	} 
	fmt.Println("person city", person.city) // 提升字段
	// 匿名结构体
	manager := struct {
		level int
		salary float64
	}{
		level: 7,
		salary: 700000.0,
	}
	fmt.Println("manager ", manager, manager.salary)
	// 结构体指针
	managerPointer := &manager
	fmt.Println("manager pointer:", managerPointer, (*managerPointer).level, managerPointer.salary)
	// 匿名字段
	var anno Anno
	anno.int = 1
	anno.string = "no"
	fmt.Println("anno", anno, anno.int, anno.string)
	// 相等
	anno2 := anno
	fmt.Println("equal", anno==anno2) // 全部字段相等
	anno2.int = 2
	fmt.Println("equal again", anno, anno2, anno==anno2)
}
