package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup){
	defer w.Done()
	fmt.Println("doing task :", id)
}

// making all task() calls parallely (concurrently)
func main(){

	var wg sync.WaitGroup



	for i:= 0; i<= 10; i++{
		wg.Add(1)
		// go routine
		go task(i, &wg)

		// // Anonymus function
		// go func (i int){
		// 	fmt.Println(i)
		// }(i) // this is function call




		// to synchronize we dont use sleep but we use wait 
		// to make all inst work befor closing main in go routine
		wg.Wait()
	}
}