package main

import(
	"fmt"
	"runtime"
)

func Print(msg string, n int){
	for i := 0; i <= n; i++{
		fmt.Printf("%s ke %d\n", msg, i)
	}
}

func main(){
	runtime.GOMAXPROCS(5)

	go Print("tes", 5)
	go Print("tas", 3)
	Print("tus", 5)

	var input string
	fmt.Scanln(&input)
}