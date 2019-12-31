package main

import "fmt"

func main() {
	a := [5]int{5, 4, 3, 1, 2}
	for i:=0; i<5; i++ {
	min := a[i]
	minp := i
	
	for j:=i+1; j<5; j++ {
		if a[j]<min {
			min = a[j]
			minp = j
			}
		}
		temp := a[minp]
		a[minp] = a[i]
		a[i] = temp
	}
	fmt.Println(a)
}

iteration1
5 4 3 1 2
4 5 3 1 2
4 3 5 1 2
4 3 1 5 2
4 3 1 2 (5)

iteration2
3 4 1 2 (5)
3 1 4 2 (5)
3 1 2 (4) (5)

iteration3
1 3 2 (4) (5)
1 2 (3) (4) (5)

iteration4
1 (2) (3) (4) (5)

(1) (2) (3) (4) (5) [sorted]
