package main
import(
f "fmt"
)
type apollo struct{
id string
axe Axe
}
type Axe struct{
rate int
fellony float32
}
func main(){
var size int;
f.Scanln(&size);
arr:=make([]apollo,size)
//{"Xtreme",Axe{15,71.24}},

for i:=0;i<size;i++ {
f.Printf("Enter the id %d:",i+1);
f.Scanln(&arr[i].id);
f.Printf("Enter %d Rate for it (in int):",i+1);
f.Scanln(&arr[i].axe.rate);
f.Printf("Enter The %d Fellony{e.g:`62.54`}:",i+1);
f.Scanln(&arr[i].axe.fellony)
}
for index,apollo:=range arr{
f.Println("The array Order " ,index+1)
f.Printf("The id: %s And rate for it is %d Fellony Price id %.2f",apollo.id,apollo.axe.rate,apollo.axe.fellony)
f.Println();
}
}