package main

import (
	"fmt"
	"sync"
	"time"
)

func count(numbers []int, threadNum string) {
	for _, number := range numbers {
		time.Sleep(50 * time.Millisecond)
		fmt.Println(threadNum, number)
	}
}

func main() {

	for {
		fmt.Printf("Press\n1 -> Worker Groups\n2 -> Channels\n3 -> Select Statement\n")
		var learn int
		fmt.Scanf("%d", &learn)

		switch learn {
		case 1:
			workerGroups()
		case 2:
			channels()
		case 3:
			selectFunc()
		default:
			fmt.Println("Wrong INPUT")
		}

		fmt.Println("Want to try again?? y/n")
		retry := ""
		fmt.Scanf("%s", &retry)
		if retry != "y" {
			break
		}

	}
}

func selectFunc() {
	/***
	Select Statements
	*/
	fmt.Println("Learning statements")

	c1 := make(chan int)
	c2 := make(chan int)
	//receiver := make(chan int)

	go func(sender chan int) {
		i := 0
		for {
			fmt.Printf("Sending data %v from channel %v \n", i, sender)
			sender <- i
			i++
			time.Sleep(1 * time.Second)
		}
	}(c1)

	go func(receiver chan int, sender chan int) {
		for {
			fmt.Printf("Received data %v from channel %v \n", <-receiver, sender)
			data := <-receiver
			sender <- data
			time.Sleep(1 * time.Second)
		}
	}(c1, c2)

	for {
		select {
		case data := <-c1:
			fmt.Println("case c1", data)
		case data := <-c2:
			fmt.Println("case c2", data)
		}
	}
}

func channels() chan int {
	fmt.Println("Learning channels")

	c1 := make(chan int)
	num := 2
	go func(sender chan int, num int) {
		for i := 0; i < num; i++ {
			fmt.Printf("Sending data %v to channel %v\n ", i, sender)
			sender <- i
		}
	}(c1, num)

	fmt.Printf("Received data %v from channel %v\n", <-c1, c1)
	fmt.Printf("Received data %v from channel %v\n", <-c1, c1)

	//If num=2, and enable below line for deadlock
	//fmt.Printf("\nReceived data %v from channel %v", <-c1, c1)

	fmt.Printf("Finished receiving data\n\n")
	return c1
}

func workerGroups() {
	/**
	Worker Groups
	*/
	wg := sync.WaitGroup{}
	fmt.Println("Lets start concurrency")

	wg.Add(1)

	go func() {
		count([]int{1, 2, 3}, "threadOne")
		wg.Done()
	}()
	count([]int{1, 2, 3}, "threadTwo")

	wg.Wait()

	/***
	Channels
	*/
}

func countTill(num int, sender chan int) {
	for i := 0; i < num; i++ {
		fmt.Printf("\nSending data %v to channel %v ", i, sender)
		sender <- i
	}
}
