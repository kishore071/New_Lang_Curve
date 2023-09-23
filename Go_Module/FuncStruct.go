package main
import (
    f "fmt"
    )

type radius struct{
    breadth float64
    height float64
}
func (r radius) getArea() float64 {
    return  3.14*r.breadth*r.height;
}
func main() {
    var examer radius;
    examer=radius{16.89,19}
    f.Println(examer.getArea())
    example:=[]radius{{10.89,78.969},}
    f.Println(example)
    f.Printf("%.2f",example[0].getArea());
//Shows the Type of the data
    f.Printf("%T %T",example,examer)
}
