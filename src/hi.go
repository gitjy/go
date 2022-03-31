package main

import (
	"fmt"
	"reflect"
)

func main() {
	//var f interface{}
	interfaceVar()
    //variable()   
    //closures()  //闭包函数
	//defers()    //延迟执行
	//deferVar()
	//触发panic
	/*set_data(20)
	fmt.Println("main  is ok")*/
	
	//functions();  //函数定义
	//hello() 
	//typeMap()    //集合
	//slice()     //切片
}

/*func printName() {
	fmt.Println(name) 
}*/


func interfaceVar() {
	var x interface{}
	getType(x)
		fmt.Print("type=") 
		fmt.Println(reflect.TypeOf(x))
	x = 2
	getType(x)
		// 声明b变量, 尝试赋值i
    var b int = x
    // getType(b)

	x = "abc"
	getType(x)
	
	x = []string{"namo", "sam", "teda"}
	getType(x)
	x = func() {

	}
	getType(x)
	 sli := []int{2, 3, 5, 7, 11, 13}

    var i interface{}
    i = sli

    g := i[1:3]
}

type nullVar interface {

}

func getType(x interface{}) {
	fmt.Printf("type=%T value=%v\n",x, x)
}


//变量别名
func variable() {
	type it int //变量别名
	var i it
	var s int
	fmt.Println(i) 
	fmt.Println(reflect.TypeOf(i)) 
	fmt.Println(reflect.TypeOf(s)) 
}

func closures() {
	name :="new name"
	pname := func () {
		fmt.Println(name)   //必包函数式是一个"内联"语句或表达式，可以不传参数，直接调用变量
	}
	pname()
}



//延迟执行
func defers() {
	defer fmt.Println("defer 1")   //直接执行语句
	defer func() {
		fmt.Println("defer 2")
	}()    //申明匿名函数且直接执行
	hello()
}

//延迟执行参数传递变量快照
func deferVar() {
	name := "go"
    defer fmt.Println("defer 1 " + name) // 输出: go
    defer func(){
	    fmt.Println("defer 2 " + name) // 输出: php
	}()
	defer func(name string){
	    fmt.Println("defer 3 " + name) // 输出: go 
	    name = "swift"
	}(name)
	defer fmt.Println("defer 4 " + name) // 输出: go

    name = "php"
    fmt.Println(name)      // 输出: defer 2
	hello()
}


//定义匿名函数
func functions() {
	hello()
	swap := func (x, y string) (string, string) {
	   return y, x
	}
	fmt.Println(reflect.TypeOf(swap))
	a, b := swap("Google", "Runoob")
   fmt.Println(a, b)
}

func hello() {
	fmt.Println("hello world!")
}



func typeMap() {
    scores := map[string]int{"english": 80, "chinese": 85}
	//fmt.Println(scores["english"])
	v := scores["english"]
	fmt.Println(v);
    math, ok := scores["english"]
    if ok {
        fmt.Printf("math 的值是: %v\n", math)
    } else {
        fmt.Println("math 不存在")
    }

    for subject := range scores {
        fmt.Printf("key: %s\n", subject)
    }

	var c map[string]string/*创建集合 */
	fmt.Printf("value=%v len=%d type=%T\n",c,len(c),c) 
	fmt.Println(reflect.TypeOf(c))
    c = make(map[string]string)
    fmt.Printf("value=%v len=%d type=%T\n", c,len(c),c)
    fmt.Println(reflect.TypeOf(c))

    /* map插入key - value对,各个国家对应的首都 */
    c [ "France" ] = "巴黎"
    c [ "Italy" ] = "罗马"
    c [ "Japan" ] = "东京"
    c [ "India " ] = "新德里"
    fmt.Println(c)

    var classCap = make(map[string]string)
    classCap["class1"] =  "Num1"
    var m map[string]string
    m = map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
    fmt.Println(m)

}


func printSlice(x []int){
   fmt.Printf("len=%d cap=%d type=%T slice=%v\n",len(x),cap(x),x, x)
}

func slice() {
	//切片和数组
	arr :=[7]int{6,5,4,3,2,1}  //数组
	fmt.Println(arr)
	//打印数组
	fmt.Println(arr[:], arr[2:],arr[:2])
	//arr =[3]int{1,2,3}  //不能重新设置的长度为3数组
	//arr =[]int{1,2,3}  //不能重新设置变量为切片类型
	arr2 :=arr  //数组
	fmt.Printf("len=%d cap=%d type=%T value=%v\n", len(arr2),cap(arr2),arr2, arr2)
	fmt.Printf("len=%d cap=%d type=%T value=%v\n", len(arr[1:3]),cap(arr[1:3]),arr[1:3], arr[1:3])

	//arr3 :=[5]string{"lisi","25","man","single"}
	slice :=arr[1:4]  //使用数组做切片
	//slice = arr3[:]    
	 
	 printSlice(slice)
	 //slice1 :=arr[1:3]  //使用数组做切片
	 //var arr3 [2]int=arr[1:3]  //使用数组片段做数组
	 //fmt.Printf("len=%d cap=%d type=%T value=%v\n", len(arr3),cap(arr3),arr3, arr3)
	 slice =[]int{1}
	 /* 同时添加多个元素 */
   slice = append(slice, 2,3,4)
    printSlice(slice)

    // 追加一个切片, ... 表示解包，不能省略
    slice = append(slice, []int{7, 8}...)
    printSlice(slice)
    slice = append([]int{0}, slice...);
    printSlice(slice)

    var numbers4 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    myslice := numbers4[4:6:8]
    myslice[1] = 10; 
    fmt.Printf("myslice为 %d, 其长度为: %d,其容量为: %d\n", myslice, len(myslice), cap(myslice))
     myslice = myslice[:4] //对数组的引用，所以会获取到后两个元素
    fmt.Printf("myslice为 %v, 其长度为: %d，其容量为: %d\n", myslice, len(myslice),cap(myslice))
}


//捕获异常
func set_data(x int) {
    defer func() {
        // recover() 可以将捕获到的panic信息打印
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()
        // 故意制造数组越界，触发panic
    var arr [10]int
    arr[x] = 88
    fmt.Println("func  is ok") //数组越界函数下面语句的无法执行
}
