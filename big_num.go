package main
import ( "fmt" )
func main(){
list:= [5]int{10,80,200,240,150}
max:=list[0]
for _,v := range(list){
if v > max{
max=v
}
}
fmt.Println("The biggest number is:",max)
}
