package main
import(
f "fmt"
)
type Rectangle struct{
base int
height int
}
func (rect Rectangle) square() int{
return rect.base*rect.height;
}
func main(){
triangle:=Rectangle{
10,
20};
var total int;
total=triangle.square();
f.Println(total);
}