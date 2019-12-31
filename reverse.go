package main
import ( "fmt" )
func main(){
str:="pavithra"
str_new:=[]byte(str)
q:=len(str_new)/2
var temp byte
for i,j := 0 ,len(str_new)-1 ; i<q ; i,j=i+1,j-1{
temp=str_new[i]
str_new[i]=str_new[j]
str_new[j]=temp
}
fmt.Println("The reverse of a string is:", string(str_new))

}
