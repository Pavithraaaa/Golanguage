package main
import ( "fmt" )
func sum(x,y int)int{
return (x+y) 
}
func diff(x ,y int)int{
return (x-y)
}
func prod(x ,y int)int{
return (x*y)
}
func division(x ,y int)int{
return (x/y)
}
func main(){
x := 10
y := 5
a:=sum(x,y)
b:=diff(x,y)
c:=prod(x,y)
d:=division(x,y)
fmt.Println("The Sum of two numbers is:",a)
fmt.Println("The diff of two numbers is:",b)
fmt.Println("The prod of two numbers is:",c)
fmt.Println("The division of two numbers is:",d)
}


