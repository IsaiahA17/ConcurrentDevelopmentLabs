//    barrier2.go - a reusable barrier implementation using channels
//    Copyright (C) 2025  Isaiah Andres

//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.

//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.

//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <https://www.gnu.org/licenses/>.
//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Isaiah Andres (C00286361@setu.ie)
// Modified On: 27/11/2025
// Description:
// A simple reusable barrier implemented using an atomic integer and unbuffered channel
// Issues:
// None I hope
//1. Change mutex to atomic variable
//2. Make it a reusable barrier
//--------------------------------------------

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, count *int32, max int32, wg *sync.WaitGroup, theChan *chan bool, theChan2 *chan bool) bool {
	for i := 0; i < 3; i++ {
		mCount := atomic.AddInt32(count, 1) //Atomic integer instead of mutex lock
		if mCount == max {                  //Each goroutine has their own count to check
			*theChan2 <- true //Last goroutine signals the previous ones to move and waits in the next line
			<-*theChan2
		} else {
			<-*theChan2       //If the goroutine isn't the last one, stay here
			*theChan2 <- true //once we get through send signal to next routine to continue including the last one
		} //end of if-else
		fmt.Println("Part A", goNum)

		mCount = atomic.AddInt32(count, -1) //Adding minus one since there's no option for subtraction
		//similar mechanism as above barrier. The above barrier is required since there's a chance that one goroutine might be too fast and output Part A while others are still on Part B
		if mCount == 0 { //last to arrive -signal others to go
			*theChan <- true
			<-*theChan
		} else { //not all here yet we wait until signal
			<-*theChan
			*theChan <- true //once we get through send signal to next routine to continue
		} //end of if-else

		fmt.Println("PartB", goNum)
	}
	wg.Done()
	return true
} //end-doStuff

func main() {
	var totalRoutines int32 = 10
	var wg sync.WaitGroup
	wg.Add(int(totalRoutines))
	//we will need some of these
	theChan := make(chan bool) //use unbuffered channel in place of semaphore
	theChan2 := make(chan bool)
	var count int32
	for i := range int(totalRoutines) { //create the go Routines here
		go doStuff(i, &count, totalRoutines, &wg, &theChan, &theChan2)
	}
	wg.Wait() //wait for everyone to finish before exiting
} //end-main
