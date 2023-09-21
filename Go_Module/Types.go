package main
import (
f "fmt"
)
type named struct{
firstname string
lastname string
age int8
}
func main(){
var name named;
f.Scanln(&name.firstname);
f.Scanln(&name.lastname);
f.Scanln(&name.age);
f.Println(name);
}