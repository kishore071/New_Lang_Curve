package main

import (
	f "fmt"
	b "bufio"
	"os"
)
func main(){
	sc:=b.NewReader(os.Stdin);
	var number int;
	f.Scanln(&number);
 	name,err:=sc.ReadString('\n');	
	f.Println("Hello");
	if err==nil {
		f.Printf("Entered Text with Number is: %s and %d",name,number);
	}
}