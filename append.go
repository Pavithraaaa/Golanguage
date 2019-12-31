package main
import ( "fmt" 
         "os"
         )
func main(){
file , err := os.OpenFile("/Users/ftt.pavithra.r/Golanguage/append.txt",os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)
if err!=nil{
fmt.Println(err)
}
write,err := file.WriteString("The end")
fmt.Println(write)
}
