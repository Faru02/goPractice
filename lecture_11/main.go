package main 

import "fmt" 
 
func main(){

type Book struct{
Title string
Author string
ISBN string 
Price float32
Pages int
}
var b1 Book
b1.Title="Golang"
b1.Author="CALEB DOXY"
b1.ISBN="978-1478355823"
b1.Price=11.43
b1.Pages=154

fmt.Println(b1)





}