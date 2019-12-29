package main
import ( "fmt"
         "regexp")
func main(){
var str string
fmt.Println("Enter an emailid")
fmt.Scanln(&str)
re := regexp.MustCompile("^[a-zA-Z0-9_]+@[a-zA-Z.]+.com$")
if re.MatchString(str)  {
 fmt.Println("The entered email is valid")
}else {
 fmt.Println("The entered email is invalid")

}
}
