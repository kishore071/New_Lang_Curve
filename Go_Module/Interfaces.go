package  main
import(
f "fmt"
)
type name interface{
printName() string
}
type users struct{
name string 
prof string
salary float64
}
type def struct{
name string
job string
timed int
salary float64
}

func (d def) printName() string{
return f.Sprintf("Name is : "+d.name+" Profession is: "+d.job+" Salary is $%v",d.salary," Per Hour,Extension Time is: %v",d.timed);
}

func (u users) printName() string{
return f.Sprintf("Name is : "+u.name+" Profession is: "+u.prof+" Salary is $%v",u.salary,"Per Hour");
}
func test(n name){
f.Println(n.printName())
}

func main(){
f.Println("Web Devloper Details");
	test(users{"Kishore","Web Dev",9865.2});
f.Println("AppDeveloper Details")
	test(def{"Kishore","App DEV",987,6789.89});
}