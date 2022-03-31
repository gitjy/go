package main

import (
    "fmt"
    "reflect"
)



func main()  {
    var phone Phone

    fmt.Println("phone type:" , reflect.TypeOf(phone)) 
    phone = new(NokiaPhone)
    fmt.Println("phone type:" , reflect.TypeOf(phone)) 
    phone.call()

    phone = NokiaPhone{}
    fmt.Println("phone type:" , reflect.TypeOf(phone)) 
    phone.call()

    phone = new(IPhone)
     fmt.Println("phone type:" , reflect.TypeOf(phone)) 
    phone.call()
    //继承
    myCom := company{
        companyName: "Tencent",
        companyAddr: "深圳市南山区",
    }
    staffInfo := staff{
        name:     "小明",
        age:      28,
        gender:   "男",
        position: "云计算开发工程师",
        company: myCom,
    }

    fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
    fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)

    // 实例化
    /*myself := Profile{name: "小明", age: 24, gender: "male"}
    // 调用函数
    myself.FmtProfile()
    myself.increase_age()
    fmt.Printf("当前年龄：%d\n", myself.age)
    mother :=Profile{name: "ada", age: 40, gender: "female"}
    myself.mother = &mother
    //指针选择器
    myself.mother.FmtProfile()
    fmt.Println(myself.mother.name, (*myself.mother).name)*/
}

type Phone interface {
    call()
}

type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
    fmt.Println("I am iPhone, I can call you!")
}


type company struct {
    companyName string
    companyAddr string
}

type staff1 struct {
    name string
    age int
    gender string
    position string
    company
}

type staff struct {
    name string
    age int
    gender string
    position string
    company
}



//定义一个名为Profile 的结构体
type Profile struct {
    name   string
    age    int
    gender string
    mother *Profile // 指针
    father *Profile // 指针
}

// 定义一个与 Profile 的绑定的方法
func (person Profile) FmtProfile() {
    fmt.Printf("名字：%s\n", person.name)
    fmt.Printf("年龄：%d\n", person.age)
    fmt.Printf("性别：%s\n", person.gender)
}

// 重点在于这个星号: *
func (person Profile) increase_age() {
    person.age += 1
}