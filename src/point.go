package main

import (
 "fmt"
)

//指针的使用
func main () {
	a := 1
	printInt(&a) //必须是地址传递

	book1 :=Books{"GO 语言", "liming", "GO语言教程", 15}
    fmt.Println(book1,&book1)
    fmt.Printf("type=%T type=%T\n", book1, &book1)
	printBook(&book1)
}

func printInt(a *int) {
	fmt.Println(a)
    fmt.Println(*a)
}

type Books struct{
	title string
	author string
	subject string
	book_id int
}

func printBook(book * Books){
     fmt.Printf("Book  title: %s\n", book.title)
     fmt.Printf("Book  author: %s\n", book.author)
     fmt.Printf("Book  subject: %s\n", book.subject)
     fmt.Printf("Book  book_id: %d\n", book.book_id)
     fmt.Printf("type=%T\n", book)
     fmt.Printf("Book  title: %s\n", (*book).title)

}

