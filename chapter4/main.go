package main

import "fmt"

//使用Printf函数进行打印

var a, b = 189, 453 //定义两个整型变量
var hello interface{}

func main() {

	runIf()     //if 测试
	runSwitch() //switch测试
	runFor()    //for测试
	runSelect() //select测试

	definedThreeDot("jack", "rose", "tom", "jerry") //定义多个参数来使用可变参数

}

func definedThreeDot(source ...string) { //定义可变参数，定义时在类型前面加上三个点
	useThreeDot(source...)     //将可变参数作为可变参数使用，使用时在定义后面加上三个点
	useThreeDotAsSlice(source) //将可变参数作为slice使用
}

func useThreeDotAsSlice(ss []string) { //定义slice来接收可变参数
	fmt.Println(ss) //直接打印slice
}

func useThreeDot(ss ...string) { //定义可变参数，定义时在类型前面加上三个点
	for index, s := range ss { //作为slice来遍历可变参数
		fmt.Printf("index : %d, value : %s \n", index, s) //index和s都作为可变参数来使用
	}
}

func runIf() {
	fmt.Printf("max value in %d, and %d is %d \n", a, b, max(a, b)) //结果： max value in 189, and 453 is 453

	//断言例子
	hello = "helloworld"

	//断言简单语法
	//value := element.(type) //type为要转换的类型

	fmt.Println(hello.(string)) //转换为string类型，能成功转换。结果：helloworld
	//fmt.Println(hello.(int))//转换为int类型，该行会报错： panic: interface conversion: interface {} is string, not int 。因为hello实际类型是string类型

	//断言高级语法
	// value, ok := element.(type) //type为要转换的类型，ok为是否成功转换，类型为bool，value为实际转换的值
	helloS, ok := hello.(string)
	if ok {
		fmt.Println("hello tranfer successfully : ", helloS) //执行该行
	} else {
		fmt.Println("hello transfer failed")
	} //结果： hello tranfer successfully :  helloworld

	//map的断言
	// value, ok := m[key] //这里的OK不再是简单的成功或者失败，理解成是否存在更合适
	var m = make(map[string]interface{}) //创建map的方式，具体make的用法后续会讲解

	m["key1"] = "value1"
	value1, ok := m["key1"]
	if ok {
		fmt.Println("map m contain 'key1' ", value1) //执行该行
	} else {
		fmt.Println("map m contain 'key1'")
	} //结果：map m contain 'key1'  value1
}

func runSwitch() {
	//switch 的语法
	grade := 10
	switch grade {
	//case code < 60://code为int类型，不能使用code < 60作为case条件
	case 10:
		fmt.Println("不及格") //执行该行
	case 70:
		fmt.Println("及格")
	default:
		fmt.Println("无效的分数")
	} //结果：不及格

	//用于类型断言
	switch hello.(type) {
	case string:
		fmt.Println("hello is string") //执行该行
	case int:
		fmt.Println("hello is int")
	default:
		fmt.Println("hello is unknown type")
	} //结果：hello is string

	switch { //直接判断case
	case a < b:
		fmt.Println("a less than b")
		fallthrough //紧接着执行下一个case，不需要进行判断
	case a > b:
		fmt.Println("a bigger than b")
	} //结果：a less than b 手动换行 a bigger than b
}

func runFor() {
	a := []int{1, 3, 9, 4, 1, 4, 6, 132, 1, 29, 43, 55, 89, 46}

	//语法一
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}

	fmt.Println(a) //结果：[1 1 1 3 4 4 6 9 29 43 46 55 89 132]

	var i = 0
	//语法二
	for i < len(a) {
		fmt.Print(a[i], "  ")
		i++
	} //结果： 1  1  1  3  4  4  6  9  29  43  46  55  89  132
	fmt.Println()

	//语法三
	i = 0
	for {
		if i < len(a) {
			fmt.Print(a[i], " ")
			i++
		} else {
			break
		}
	} //结果： 1 1 1 3 4 4 6 9 29 43 46 55 89 132
	fmt.Println()

	//语法四
	for index, value := range a {
		fmt.Printf("index: %d, value: %d \n", index, value)
	}
	/*
	结果：
	index: 0, value: 1
	index: 1, value: 1
	index: 2, value: 1
	index: 3, value: 3
	index: 4, value: 4
	index: 5, value: 4
	index: 6, value: 6
	index: 7, value: 9
	index: 8, value: 29
	index: 9, value: 43
	index: 10, value: 46
	index: 11, value: 55
	index: 12, value: 89
	index: 13, value: 132
	 */

	m := map[string]string{}
	m["hello"] = "world"
	m["hey"] = "bye"

	for key, value := range m {
		fmt.Printf("key: %s, value: %s \n", key, value)
	}
	/*
	结果：
	key: hello, value: world
	key: hey, value: bye
	 */
}

func runSelect() {
	fmt.Println("select")
	c := make(chan int, 1)
	select {
	case c <- 1:
		fmt.Println("push into channel")
	case <-c:
		fmt.Println("get from channel")
	default:
		fmt.Println("default")
	}
}

func max(a, b int) (max int) {
	if a > b {
		max = a
	} else if a == b {
		max = a
	} else {
		max = b
	}

	return
}
