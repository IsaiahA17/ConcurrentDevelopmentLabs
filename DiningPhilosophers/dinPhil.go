// Dining Philosophers Template Code
// Author: Joseph Kehoe
// Created: 21/10/24
// Modified By: Isaiah Andres (C00286361@setu.ie)
// Date Modified: 27/11/25

// dinPhil.go is a solution to the Dining Philosophers problem
// Copyright (C) 2025  Isaiah Andres

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

func think(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second)                  //wait random time amount
	fmt.Println("Phil: ", index, "was thinking") //Simulating Work
}

func eat(index int) {
	var X time.Duration
	X = time.Duration(rand.IntN(5))
	time.Sleep(X * time.Second) //wait random time amount
	fmt.Println("Phil: ", index, "was eating")
}

func getForks(index int, forks map[int]chan bool) {
	if index == 0 { //If the first goroutine/index 0 grab right fork and then left fork, always leaving one free for the other philosopher to break the circular wait that happens if all philosophers grab the left fork
		forks[(index+1)%5] <- true //Grabbing fork to the right and blocking if sending to a channel that already has a value
		forks[index] <- true       //Grabbing fork to the left
	} else { //If not the first routine, grab left and then right
		forks[index] <- true
		forks[(index+1)%5] <- true
	}
}

func putForks(index int, forks map[int]chan bool) {
	<-forks[index] //Put left and right fork back, allowing waiting goroutines to continue and eat thus achieving synchronisation while the current one moves on to think and then wait
	<-forks[(index+1)%5]
}

func doPhilStuff(index int, wg *sync.WaitGroup, forks map[int]chan bool) {
	for {
		think(index)
		getForks(index, forks)
		eat(index)
		putForks(index, forks)
	}
}

func main() {
	var wg sync.WaitGroup //Creating Waitgroup
	philCount := 5        //Adding five goroutines to the WaitGroup
	wg.Add(philCount)

	forks := make(map[int]chan bool) //Creating map of boolean channels that can be accessed by an integer index
	for k := range philCount {
		forks[k] = make(chan bool, 1) //Assigning a boolean channel of size 1 to each fork
	} //set up forks
	for N := range philCount {
		go doPhilStuff(N, &wg, forks)
	} //start philosophers
	wg.Wait() //wait here until everyone (10 go routines) is done

} //main
