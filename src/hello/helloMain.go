package main //可运行才能是main包
import (
	"fmt"
	"math"
	_ "structs/exportStruct" // _ 包引入但是不用，不加_会报错，go对没用到的包报错，减少不必要的引入，减少编译时间和包大小
)
/*
 * 返回单个值
 */
func calculatBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

/*
 * 返回多个值
 */
func testReturnMultiValue(a, b int)(int, int) {
	var sum = a + b
	var minus = a - b
	return sum, minus
}

/*
 * 命名返回值
 */
func namedMultiReturn(a, b int)(sum, diff int) {
	sum = a + b
	diff = a - b
	return
}

/*
 * 函数传指针，修改指向的值
 */
func changePointerValue(pointer *int) {
	*pointer += 1000
}
// 大小为3的数组
func changeArrayValue(arr *[3]int) {
	(*arr)[0] = -1
	arr[1] = -2
}
// 传切片
func changeArrayValueBySlice(arr []int) {
	arr[0] = -11
}

type Employee struct {
	name, email string
	age int
}
/*
 * 方法定义
 */
func (e Employee) displayEmployee() {
	fmt.Println("employee", e.name, e.age)
}
/*
 * 指针作为接收器
 */
func (e *Employee) changeEmployeeName(newName string) {
	//(*e).name = newName
	e.name = newName
}
/*
 * 非结构体的方法
 */
type myInt int
func (a myInt) add(b myInt) myInt {
	return a + b
}

func main() {
	/*
	 * 变量声明和赋值学习
	 */
	var age0 int // 变量声明
	age0 = 12
	fmt.Println("age0 is", age0)

	var age int = 25 // 变量声明并初始化
	var width, height = 1, 2 //类型推断
	fmt.Println("hello world, my age is", age)
	fmt.Println("width is", width, "height is", height)
	age = 26
	fmt.Println("my age is", age, "now")
	// 不同类型变量
	var (
		name = "xieyizun"
		maxage = 100
	)
	fmt.Println("my name is", name, ", maxage is", maxage, ", height", height)

	// := 左边至少需要一个没有声明的变量
	a, b := "xieyizun", 100 // 变量简短声明
	// 重新赋值, 强类型 a已经是string类型，不能重新赋值整型
	a, b = "1", 2
	c, d := 1.1, 2.0
	fmt.Println("c d larger is", math.Max(c, d))
	fmt.Println("a is", a, "b is", b)

	/*
	 * 变量类型学习
	 */
	 var f float32 = 4.5
	 fmt.Printf("f type is %T, value is %f \n", f, f)
	 i, j := 1, 2.4
	 sum := float64(i)+j
	 fmt.Println("sum is", sum)

	 /*
	  * 常量学习，常量是无类型的
	  */
	const cc = 12
	const c_str = "hello world"
	fmt.Println("const cc is", cc, c_str)

	/*
	 * 函数
	 */
	var totalPrice = calculatBill(12, 4)
	fmt.Println("total price is", totalPrice)
	s, m := testReturnMultiValue(10, 5)
	fmt.Printf("sum is %d, minus is %d\n", s, m)
	s2, _ := namedMultiReturn(10, 5)
	fmt.Println(s2)

	/*
	 * if else
	 */
	 if b=11; b % 2 == 0 {
		 fmt.Println("a is 偶数")
	 } else if b % 2 != 0 {
		 fmt.Printf("a %d is 奇数\n", b)
	 } else { // else 语句应该在 if 语句的大括号 } 之后的同一行中。如果不是，编译器会不通过。
		 fmt.Println("none")
	 }

	 /*
	  * for循环
	  */
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		if i == 9 {
			break
		}
		fmt.Println("for i=", i)
	}
	i = 0
	for ; i < 10; i++ {
		fmt.Println("for again is", i)
	}
	for j, i := 0, 0; i < 10; i, j = i+1, j+1 {
		fmt.Println("i, j ", i, j)
	}
	// // 无限循环
	// for {
	// }

	/*
	 * switch
	 */ 
	switch ii := 1; {
	 	case ii < 1 || ii < 2 || ii < 3:
			fmt.Println("<3")
			fallthrough // 继续下个case 不退出
		case ii < 4:
			fmt.Println("<4 too")
		default:
			fmt.Println("default")
	}
	jj := 1
	switch jj {
	case 1, 2:
		fmt.Println("jj < 3", jj)
	}

	/*
	 * 数组
	 */
	 var arr1 [3]int
	 arr1[0] = 1
	 arr1[1] = 2
	 arr1[2] = 3
	 arr2 := [3]int{4,5,6}
	 arr3 := [...]string{"7","8","9"}
	 arr4 := arr3 // 值引用，改变arr4，arr3不变
	 arr4[0] = "hello"
	 arr5 := [2][1]int{{1}, {2}} // 多维数组，两行一列
	 fmt.Println("arrs is", arr1, arr2, arr3)
	 //遍历
	 for i := 1; i < len(arr1); i++ {
		 fmt.Println("arr1 value: ", arr1[i])
	 }
	 for i, v := range arr3 {
		 fmt.Println("arr1 again index, value: ", i, v)
	 }
	 for i, row := range arr5 {
		 for _, col := range row {
			 fmt.Printf("row %d, col value: %d\n", i, col)
		 }
	 }
	 // 分片：引用底层数组，修改也会修改底层数组
	 slice1 := arr1[1:3]
	 slice1[0] = 100
	 //var slice2 []int = arr1[1:3]
	 for _, v := range slice1 {
		 fmt.Println("slice value ", v)
	 }
	 // slice1底层数组也修改了,slice5为拷贝，不会修改且arr1不会因为slice5的存在而无法被回收
	 slice5 := make([]int, len(slice1))
	 // 拷贝
	 //copy(slice5, slice1)
	 copy(slice5, arr1[1:3])
	 fmt.Println("slice5", slice5)
	 slice5[0]=1000
	 for _, v := range arr1 {
		fmt.Println("arr1 value ", v)
	 }
	 // 创建切片
	 slice3 := []int{4,5,6}
	 slice4 := make([]int, 2, 6) // 2为长度，6为容量。长度超过6时，容量翻倍
	 slice3[0] = 1 // 修改
	 slice3 = append(slice3, 7) // 追加，数组不支持添加
	 slice4 = append(slice4, slice3...)
	 slice4 = append(slice4, 100)
	 fmt.Println("slice3, slice4", slice3, slice4, len(slice4), cap(slice4))
	 // 切片作为函数参数，在函数内被修改了，切片自身也会被修改
	 // 可变函数：函数的可变参数是切片类型，即会将传入的可变参数转为切片类型。如func a(num ...int), 调用：a(slice1...)，
	 // 此时如果在函数a里面修改了slice1，则外面的slice也会修改

	 /*
	  * map字典：personSalary := make(map[string]int)
	  */
	// 创建
	var map1 map[string]int // 此时为nil
	map1 = make(map[string]int)
	map1["key"] = 1
	map2 := make(map[string]string)
	map2["key"] = "value"
	map3 := map[string]int {
		"key1": 1,
		"key2": 2,
	}
	// 获取
	value, ok := map3["key1"] // ok是否存在key1
	fmt.Println("map1,2,3: ", map1, map2, map3, len(map3))
	fmt.Println("map1 key value", map1["key"])
	fmt.Println("map3 value ok", value, ok)
	// 遍历,不保证每次执行程序获取的元素顺序相同
	for key, value := range map3 {
		fmt.Println("遍历 map3: ", key, value)
	}
	// 删除
	delete(map3, "key1")
	value, ok = map3["key1"] // ok是否存在key1
	fmt.Println("map3 again value ok", value, ok)
	// map是引用类型
	map4 := map3
	map4["key1"] = 4
	value, ok = map3["key1"] // ok是否存在key1
	fmt.Println("map3 again again value ok", value, ok)
	// map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil，即不能map3==map4
	fmt.Println("map4 is nil", map4 == nil)

	/*
	* 字符串，不可变，如str := "34,54", str[0]='1'是不行的,若要修改，需要转为一个rune切片：[]rune(str)
	*/
	// Go 中的字符串是兼容 Unicode 编码的，并且使用 UTF-8 进行编码。
	// %x：十六进制打印；%c：字符打印
	// rune：go内建类型，int32别称
	str1 := "Hello world"
	chars2 := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9} // 用切片构造字符串
	str2 := string(chars2)
	fmt.Println("s1, s2: ", str1, str2)
	// rune使用
	chars3 := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
    str3 := string(chars3)
    fmt.Println(str3)
	for index, rune := range str3 {
		fmt.Printf("str char in %d, %c\n", index, rune)
	}
	// 长度：len返回字节个数

	/*
	* 指针：存储变量内存地址的变量
	*/
	point1Str := "hello world"
	pointer1 := &point1Str
	point2int := 3444
	var pointer2 *int = &point2int // 指针变量的类型为 *T，该指针指向一个 T 类型的变量
	fmt.Println("point1 point2", pointer1, pointer2)
	// 指针解引用
	fmt.Println("point1 value", *pointer1)
	var zero *int // 指针零值为nil
	if zero == nil {
		fmt.Println("pointer zero is nil")
	}
	// 函数传指针
	funcPointer := 1
	changePointerValue(&funcPointer)
	fmt.Println("changePointerValue", funcPointer)
	// 数组，传指针和切片都会修改，不过切片更简洁
	arr := [3]int{1,2,3}
	changeArrayValue(&arr) //数组指针
	fmt.Println("changeArrayValue", arr)
	changeArrayValueBySlice(arr[:]) // 切片
	fmt.Println("changeArrayValueBySlice", arr)
	//go不支持指针运算
	// arrPointer := &arr
	// arrPointer++ // 报错

	/*
	* struct结构体
	*/
	employee2 := Employee{
		name: "xieyizun",
		email: "xxx",
		age: 25,
	}

	/*
	* 方法，注意不是函数，
	* func (t Type) methodName(parameter list) {}， Type为方法接收器
	*/
	employee2.changeEmployeeName("xieyizun2")
	(&employee2).changeEmployeeName("xieyizun3")
	employee2.displayEmployee()
	aInt := myInt(1)
	bInt := myInt(2)
	sumInt := aInt.add(bInt)
	fmt.Println(sumInt)
	
}
