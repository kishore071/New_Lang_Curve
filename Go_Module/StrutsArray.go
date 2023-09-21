package main
import(
f "fmt"
)
type Array struct{
name string
cost int
}
func main(){
array:=[]Array{
{"Post",100},
{"Put",56},
{"Get",90},
}
f.Println("Length :",len(array),"Capacity :",cap(array))
f.Println(array);
}