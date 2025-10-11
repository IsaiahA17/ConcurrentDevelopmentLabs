//Barrier.go Template Code
//Copyright (C) 2024 Dr. Joseph Kehoe

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

//--------------------------------------------
// Author: Joseph Kehoe (Joseph.Kehoe@setu.ie)
// Created on 30/9/2024
// Modified by:
// Issues:
// The barrier is not implemented!
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
	mCount := atomic.AddInt32(count, 1)
	if mCount == 10 {
		*barrier <- struct{}{}
	}
	<-*barrier
	*barrier <- struct{}{}
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
