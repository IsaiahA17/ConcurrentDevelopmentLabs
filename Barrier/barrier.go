// Barrier Code
// Author: Isaiah Andres
// Created: 29/09/25

// barrier.go is an implementation of a barrier with a buffered channel and an atomic integer
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
//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by: Isaiah Andres (C00286361@setu.ie)
// Issues:
// Hopefully None
//--------------------------------------------

package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"
)

var count int32

// Place a barrier in this function --use Mutex's and Semaphores
func doStuff(goNum int, wg *sync.WaitGroup, count *int32, barrier *chan struct{}) bool {
	time.Sleep(time.Second)
	fmt.Println("Part A", goNum)
	mCount := atomic.AddInt32(count, 1) //Add one to the number of goroutines entered, atomic integers ensure only one goroutine can increment and is faster than a mutex lock
	if mCount == 10 {                   //mCount is a variable local to this function and will be used in the if statement so they don't share a variable, making it more threadsafe
		*barrier <- struct{}{} //Last goroutine signals the previous ones to move and waits in the next line
	}
	<-*barrier             //Taking from an empty channel blocks goroutines
	*barrier <- struct{}{} //Unblock other goroutines to move on to print Part B
	fmt.Println("PartB", goNum)
	wg.Done()
	return true
}

func main() {
	totalRoutines := 10
	var wg sync.WaitGroup
	wg.Add(totalRoutines)
	//we will need some of these
	ctx := context.TODO()
	var theLock sync.Mutex
	sem := semaphore.NewWeighted(int64(totalRoutines))
	theLock.Lock()
	sem.Acquire(ctx, 1)
	barrier := make(chan struct{}, 1)
	for i := range totalRoutines { //create the go Routines here
		go doStuff(i, &wg, &count, &barrier)
	}
	sem.Release(1)
	theLock.Unlock()

	wg.Wait() //wait for everyone to finish before exiting
}
