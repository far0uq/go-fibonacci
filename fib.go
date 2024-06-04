package main

import "fmt"

// Function to compute the fibonacci series number at a given index
func computeFib(num int) int {
    fib := []int{0,1};
    for i:=2; i<=num; i++{
        fib = append(fib, fib[i-2]+fib[i-1]);
    }
    fmt.Println(fib);
    return fib[num];
}

func main(){
    fmt.Println(computeFib(5));
}