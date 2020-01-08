package main

import "fmt"

func main() {
	a := []int{1, 5, 3, 4, 2}
	for i := 0;i < 5;i++ {       
		for j:=i+1; j < 5; j++ { 
			if a[i] > a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}

