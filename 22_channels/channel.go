package main

import (
	"fmt"
	"sync"
	// "math/rand"
)

func processNum(numChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range numChan {
		fmt.Println("processed num", num)
	}

}

func sum(result chan int, num1 int, num2 int) {
	numresult := num1 + num2
	result <- numresult

}

func task(done chan bool) {
	defer func() {
		done <- true
	}()
	fmt.Println("Processing ...")
}

// buffered channel :-
// limited amount of data can be send without blocking

// que system implementation
func emailSender(emailcha chan string, done chan bool) {

	defer func() {
		done <- true
	}()

	for email := range emailcha {
		fmt.Println("sending emial to :", email)
	}
}

func main() {
	// var wg sync.WaitGroup
	// msgchannel := make(chan int)

	// // msgchannel <- "ping"  //sending data in channel      // blocking
	// // msg := <- msgchannel
	// // fmt.Println(msg)



	// wg.Add(1)
	// go processNum(msgchannel, &wg)

	// for{
	// 	msgchannel <- rand.Intn(100)
	// }

	// wg.Wait()

	result := make(chan int)


	
	go sum(result, 4, 5)

	res := <-result // blocking
	fmt.Println(res)

	// done := make(chan bool)

	// go task(done)

	// a := <- done // blocking

	// fmt.Println(a)

	// buffered channel
	emailcha := make(chan string, 100)
	// emailcha <- "uday@123.com"
	// emailcha <- "maharana@.com"

	// using done as sync mech insted of wait and sleep
	done := make(chan bool)

	go emailSender(emailcha, done)

	for i := 0; i < 110; i++ {
		// inf loop bec it is buffered channel and will be blocked
		// below done will wait and wait only
		emailcha <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending")

	// fmt.Println(<-emailcha)
	// fmt.Println(<-emailcha)

	// closing channel
	close(emailcha)

	<-done
}
