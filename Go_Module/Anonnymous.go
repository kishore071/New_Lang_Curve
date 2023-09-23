package main
import(
f "fmt"
)
type name struct{
first string
last string
}

//Anonnymous.go

type studdet struct{
name
regno int
dob string
}
func main(){
f.Println("How Many Details u wanna enter")
var size int;
f.Scanln(&size);
details:=make([]studdet,size)
for i:=0; i<size; i++ {
f.Println("The First Name of the Student is");

f.Scanln(&details[i].first);
f.Println("The Last Name of the Student is");

f.Scanln(&details[i].last);
f.Println("Student Register No:");
f.Scanln(&details[i].regno);
f.Println("Date Of Birth Of the Student");
f.Scanln(&details[i].dob)
}
for i:=0; i<size; i++ {
f.Println("The Full Name of the Student is ",details[i].first+" "+details[i].last," and Register No is:",details[i].regno,"Date of Birth is "+details[i].dob)
}
}