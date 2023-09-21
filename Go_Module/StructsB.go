package main
import (
f "fmt"
)
type Rechange struct{
carrierName string
pack int32
status bool
}
//Pointer Concept Being Implemented
func (rec *Rechange) sucess(newStatus bool){
rec.status=newStatus;
}
func main(){
stud:=Rechange{
"Airtel",
666,
false,
}
f.Println("Before Recharge",stud);
f.Println("After Recharge:")
stud.sucess(true);
f.Print(stud);
}