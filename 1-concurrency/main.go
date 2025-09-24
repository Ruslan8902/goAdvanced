package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ch_1 := make(chan int)
	ch_2 := make(chan int)
	for i := 0; i < 10; i++ {
		go generate_random_n_to_chan(ch_1)
		go pow_two_of_n_from_chan_to_chan(ch_1, ch_2)
	}

	for i := 0; i < 10; i++ {
		n := <-ch_2
		fmt.Print(n, " ")
	}
}

func generate_random_n_to_chan(c chan int) {
	c <- rand.Intn(101)
}

func pow_two_of_n_from_chan_to_chan(c_from chan int, c_to chan int) {
	n := <-c_from
	c_to <- (n * n)
}
